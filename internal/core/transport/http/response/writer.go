package core_http_response

import "net/http"

var (
	StatisCodeUninitialazed = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatisCodeUninitialazed,
	}
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statusCode = statusCode
}

func (rw *ResponseWriter) GetSasusCode() int {
	if rw.statusCode == StatisCodeUninitialazed {
		return http.StatusOK
	}
	return rw.statusCode
}
