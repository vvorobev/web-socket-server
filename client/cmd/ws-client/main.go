package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"ws-client/internal/app"
	"ws-client/internal/config"
)

func main() {
	config := config.MustLoad()
	application := app.New(config.HTTP.Port)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		application.HTTPServer.MustRun()
	}()
	<-done
	fmt.Println("Server Stopped")
}
