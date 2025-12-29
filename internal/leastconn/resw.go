package leastconn

import (
	"net/http"
	"sync"
)

type ContingResponseWriter struct {
	http.ResponseWriter
	OnClose func()
	once    sync.Once
}

func (w *ContingResponseWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}

func (w *ContingResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *ContingResponseWriter) Finish() {
	w.once.Do(w.OnClose)
}
