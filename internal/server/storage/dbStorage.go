package storage

import (
	"context"
	"database/sql"
	"github.com/stranik28/MetricsCollector/internal/server/models"
)

func createTables(db *sql.DB) error {
	sqlStatement := `
CREATE TABLE IF NOT EXISTS counter (
    name VARCHAR(64),
    value NUMERIC,
    PRIMARY KEY (name)
);
CREATE TABLE IF NOT EXISTS gauge (
    name VARCHAR(64),
    value double precision,
    PRIMARY KEY (name)
);
`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}

func LoadMetricsFromDB(c context.Context, db *sql.DB) (map[string]Metric, error) {
	err := createTables(db)
	if err != nil {
		return nil, err
	}
	metrics := make(map[string]Metric)

	rows, err := db.QueryContext(c, "SELECT name, value FROM counter;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var value int64
		if err := rows.Scan(&name, &value); err != nil {
			return nil, err
		}
		metrics[name] = Metric{Counter: value}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	// Запрос для получения измерений
	rows, err = db.QueryContext(c, "SELECT name, value FROM gauge;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var value float64
		if err := rows.Scan(&name, &value); err != nil {
			return nil, err
		}
		metrics[name] = Metric{Gauge: value}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return metrics, nil
}

func GetMetricByName(c context.Context, db *sql.DB, metricName string, metricType string) (Metric, error) {
	var metric Metric
	var preReq *sql.Stmt
	var err error
	err = createTables(db)
	if err != nil {
		return metric, err
	}
	if metricType == "gauge" {
		preReq, err = db.PrepareContext(c, "SELECT * FROM gauge WHERE name=$1;")
		if err != nil {
			return metric, err
		}
	} else if metricType == "counter" {
		preReq, err = db.PrepareContext(c, "SELECT * FROM counter WHERE name=$1;")
		if err != nil {
			return metric, err
		}
	} else {
		err := ErrorMetricsNotFound
		return metric, err
	}

	rows, err := preReq.QueryContext(c, metricName)
	if err != nil {
		return metric, err
	}
	defer rows.Close()

	for rows.Next() {
		if metricType == "counter" {
			var name string
			var value int64
			if err := rows.Scan(&name, &value); err != nil {
				return metric, err
			}
			metric.Counter = value
			return metric, nil
		}
		var name string
		var value float64
		if err := rows.Scan(&name, &value); err != nil {
			return metric, err
		}
		metric.Gauge = value
	}
	if err := rows.Err(); err != nil {
		return metric, err
	}
	return metric, nil
}

func InsertMetric(c context.Context, db *sql.DB, metrics []models.Metrics) error {
	var err error
	err = createTables(db)
	if err != nil {
		return err
	}

	if len(metrics) == 0 {
		return nil
	}

	tx, err := db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	for _, val := range metrics {
		if val.Delta != nil {
			_, err = tx.ExecContext(c, "INSERT INTO counter (name, value) VALUES ($1, $2) "+
				"ON CONFLICT (name) DO UPDATE SET value = counter.value + EXCLUDED.value;", val.ID, val.Delta)
		} else {
			_, err = tx.ExecContext(c, "INSERT INTO gauge (name, value) VALUES ($1, $2) "+
				"ON CONFLICT (name) DO UPDATE SET value = EXCLUDED.value;", val.ID, val.Value)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
