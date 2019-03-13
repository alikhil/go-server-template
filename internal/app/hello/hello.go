package hello

import (
	"net/http"
)

// HelloHandler - simple http handler
func (ctx *Context) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello!"))
}
