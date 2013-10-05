package pagecache

import (
	"net/http"
	"sync"
	"time"
)

type responseWriter struct {
	rw     http.ResponseWriter
	header http.Header
	buff   []byte
}

func (w *responseWriter) Header() http.Header {
	return w.rw.Header()
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.buff = b
	return w.rw.Write(b)
}

func (w *responseWriter) WriteHeader(status int) {
	w.rw.WriteHeader(status)
}

func CacheHandlerFunc(handler func(w http.ResponseWriter, r *http.Request), expire time.Duration) http.Handler {
	return CacheHandler(http.HandlerFunc(handler), expire)
}

func CacheHandler(handler http.Handler, expire time.Duration) http.Handler {
	cw := new(responseWriter)
	limit := make(map[string]time.Time)
	buff := make(map[string][]byte)
	lock := sync.Mutex{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lock.Lock()
		defer lock.Unlock()

		now := time.Now()
		key := r.Method + r.RequestURI
		if l, ok := limit[key]; !ok || l.Before(now) {
			limit[key] = now.Add(expire)
			cw.rw = w
			handler.ServeHTTP(cw, r)
			buff[key] = cw.buff
		} else {
			w.Write(buff[key])
		}
	})
}
