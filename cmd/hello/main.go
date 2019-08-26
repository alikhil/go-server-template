package main

import (
	"context"
	"github.com/alikhil/go-server-template/internal/app/hello"
	"github.com/alikhil/go-server-template/pkg/utils"
	"os"
)

var log = &utils.Logger{MinLogLevel: utils.Debug}

func main() {

	var config = hello.GetConfig()
	var ctx, err = hello.GetContext(config)

	if err != nil {
		log.Fatal("failed to get context - %s", err)
	}

	var server = hello.NewServer(ctx)

	utils.RegisterGracefulShutdown(func(s os.Signal) {
		log.Info("received %v signal. exiting.", s.String())
		var err = server.Shutdown(context.Background())
		if err != nil {
			log.Error("failed to shutdown http server - %v", err)
		}
	})

	log.Info("starting application in address - %v", config.ServeAddress)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("%v", err)
	}
}
