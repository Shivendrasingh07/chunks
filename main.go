package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"example.com/m/serverProvider"

	"github.com/sirupsen/logrus"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := serverProvider.SrvInit()

	go srv.Start()

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	//})

	fmt.Println("running")
	<-done
	logrus.Info("Graceful shutdown")
	srv.Stop()

}
