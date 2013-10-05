package pagecache

import (
	"bytes"
	"net/http"
	"time"
)

type responseWriter struct {
	rw   http.ResponseWriter
	buff *bytes.Buffer
}

func (w *responseWriter) Header() http.Header {
	return w.rw.Header()
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.buff = new(bytes.Buffer)
	w.buff.Write(b)
	return w.rw.Write(b)
}

func (w *responseWriter) WriteHeader(status int) {
	w.rw.WriteHeader(status)
}

func CacheHanlder(handler http.Handler, expire time.Duration) http.Handler {
	cw := new(responseWriter)
	var limit *time.Time
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		if limit == nil || (*limit).Before(now) {
			l := now.Add(expire)
			limit = &l
			cw.rw = w
			handler.ServeHTTP(cw, r)
		} else {
			w.Write(cw.buff.Bytes())
		}
	})
}
