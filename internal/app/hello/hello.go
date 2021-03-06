package hello

import (
	"net/http"
)

// HelloHandler - simple http handler
func (ctx *Context) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hi!"))
}

// ReadinessHealthcheckHandler - dumb healthcheck method
func (ctx *Context) ReadinessHealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok!"))
}

// LivenessHealthcheckHandler - dumb healthcheck method
func (ctx *Context) LivenessHealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok!"))
}
