package main

import (
	"context"
	"go_restapi/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	serverDoneChan := make(chan os.Signal, 1)
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":5000")

	
	go func() {
		err := srv.ListenAndServe()

		if err != nil {
			panic(err)
		}

	}()

	log.Println("Server started")

	<-serverDoneChan
	srv.Shutdown(ctx)

	log.Println("Server stopped")
}
