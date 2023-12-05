package writer

import "net/http"

// CustomResponseWriter is a custom ResponseWriter that intercepts the status code
type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Msg        []byte
}

// WriteHeader intercepts the status code and stores it
func (w *CustomResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// WriteHeader intercepts the bytes written and stores them
func (w *CustomResponseWriter) Write(msg []byte) (int, error) {
	w.Msg = msg
	return w.ResponseWriter.Write(msg)
}
