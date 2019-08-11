package main

import "github.com/rafaelcam/go-transaction-service/cmd/servid"

func main() {
	server := servid.NewApp()
	server.Start()
}