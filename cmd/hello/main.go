package main

import (
	"github.com/alikhil/go-server-template/internal/app/hello"
	"github.com/alikhil/go-server-template/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var log = &utils.Logger{MinLogLevel: utils.Debug}

func main() {

	var config = hello.GetConfig()
	var ctx, err = hello.GetContext(config)

	if err != nil {
		log.Fatal("failed to get context - %s", err)
	}

	var router = mux.NewRouter()
	router.HandleFunc("/", ctx.HelloHandler).Methods("GET")

	utils.RegisterGracefulShutdown(func(s os.Signal) {
		log.Info("received %v signal. exiting.", s.String())
	})

	log.Info("starting application in address - %v", config.ServeAddress)
	log.Fatal("%v", http.ListenAndServe(config.ServeAddress, router))
}
