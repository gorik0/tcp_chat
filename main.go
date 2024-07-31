package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"tcp/chat"
)

func main() {

	panamaChan := make(chan os.Signal)
	signal.Notify(panamaChan, os.Interrupt, syscall.SIGTERM)
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server := chat.NewServer(listen)

	go func() {

		log.Println("Starting server...")
		err = server.Run()
		if err != nil {
			log.Println("Error starting server:", err)
			os.Exit(1)
		}

	}()
	<-panamaChan
	//	:: launch server
}
