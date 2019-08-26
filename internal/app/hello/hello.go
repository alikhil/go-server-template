package hello

import (
	"net/http"
)

// HelloHandler - simple http handler
func (ctx *Context) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello!"))
}

// ReadinessHealthcheckHandler - dumb healthcheck method
func (ctx *Context) ReadinessHealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello!"))
}

// LivenessHealthcheckHandler - dumb healthcheck method
func (ctx *Context) LivenessHealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello!"))
}
