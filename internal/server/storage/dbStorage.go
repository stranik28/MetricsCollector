package storage

import (
	"database/sql"
	"fmt"
)

func createTables(db *sql.DB) error {
	sqlStatement := `
CREATE TABLE IF NOT EXISTS counter (
    name VARCHAR(64),
    value integer,
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

func saveMetricsToDB(db *sql.DB) {
	err := createTables(db)
	if err != nil {
		panic(err)
	}
	sqlStatementDelete := "DELETE FROM gauge; DELETE FROM counter;"
	_, err = db.Exec(sqlStatementDelete)
	if err != nil {
		panic(err)
	}

	sqlStrCounter := "INSERT INTO counter (name, value) VALUES "
	sqlStrGauge := "INSERT INTO gauge (name, value) VALUES "
	var countVal []any
	var gaugeVal []any
	metrics, err := GetAll()
	if err != nil {
		panic(err)
	}
	countCount := 0
	countGauge := 0
	if len(metrics) == 0 {
		return
	}
	for key, val := range metrics {
		if val.Counter != 0 {
			sqlStrCounter += fmt.Sprintf(" ($%d, $%d), ", countCount+1, countCount+2)
			countVal = append(countVal, key, val.Counter)
			countCount += 2
		} else {
			sqlStrGauge += fmt.Sprintf(" ($%d, $%d), ", countGauge+1, countGauge+2)
			gaugeVal = append(gaugeVal, key, val.Gauge)
			countGauge += 2
		}
	}

	sqlStrCounter = sqlStrCounter[:len(sqlStrCounter)-2] + ";"
	sqlStrGauge = sqlStrGauge[:len(sqlStrGauge)-2] + ";"

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	if countCount != 0 {
		stmt, _ := db.Prepare(sqlStrCounter)
		_, err := stmt.Exec(countVal...)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}
	if countGauge != 0 {
		stm, _ := db.Prepare(sqlStrGauge)
		_, err := stm.Exec(gaugeVal...)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

}

func loadMetricsFromDB(db *sql.DB) (map[string]Metric, error) {
	metrics := make(map[string]Metric)

	rows, err := db.Query("SELECT name, value FROM counter")
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
	rows, err = db.Query("SELECT name, value FROM gauge")
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
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return metrics, nil
}
