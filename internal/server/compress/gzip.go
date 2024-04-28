package compress

import (
	"bufio"
	"compress/gzip"
	"io"
	"net"
	"net/http"
)

// CompressWriter реализует интерфейс http.ResponseWriter и позволяет прозрачно для сервера
// сжимать передаваемые данные и выставлять правильные HTTP-заголовки
type CompressWriter struct {
	w  http.ResponseWriter
	zw *gzip.Writer
}

func (c *CompressWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CompressWriter) Flush() {
	//TODO implement me
	panic("implement me")
}

func (c *CompressWriter) CloseNotify() <-chan bool {
	//TODO implement me
	panic("implement me")
}

func (c *CompressWriter) Status() int {
	//TODO implement me
	panic("implement me")
}

func (c *CompressWriter) Size() int {
	//TODO implement me
	panic("implement me")
}

func (c *CompressWriter) WriteString(s string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CompressWriter) Written() bool {
	//TODO implement me
	panic("implement me")
}

func (c *CompressWriter) WriteHeaderNow() {
	//TODO implement me
	panic("implement me")
}

func (c *CompressWriter) Pusher() http.Pusher {
	//TODO implement me
	panic("implement me")
}

func NewCompressWriter(w http.ResponseWriter) *CompressWriter {
	return &CompressWriter{
		w:  w,
		zw: gzip.NewWriter(w),
	}
}

func (c *CompressWriter) Header() http.Header {
	return c.w.Header()
}

func (c *CompressWriter) Write(p []byte) (int, error) {
	return c.zw.Write(p)
}

func (c *CompressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		c.w.Header().Set("Content-Encoding", "gzip")
	}
	c.w.WriteHeader(statusCode)
}

// Close закрывает gzip.Writer и досылает все данные из буфера.
func (c *CompressWriter) Close() error {
	return c.zw.Close()
}

// compressReader реализует интерфейс io.ReadCloser и позволяет прозрачно для сервера
// декомпрессировать получаемые от клиента данные
type compressReader struct {
	r  io.ReadCloser
	zr *gzip.Reader
}

func NewCompressReader(r io.ReadCloser) (*compressReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		r:  r,
		zr: zr,
	}, nil
}

func (c compressReader) Read(p []byte) (n int, err error) {
	return c.zr.Read(p)
}

func (c *compressReader) Close() error {
	if err := c.r.Close(); err != nil {
		return err
	}
	return c.zr.Close()
}
