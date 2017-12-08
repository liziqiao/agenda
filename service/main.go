package main

import "github.com/HinanawiTenshi/agenda/service/service"

func main() {
	port := ":8080"

	server := service.NewServer()
	server.Run(port)
}
