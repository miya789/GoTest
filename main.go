package main

import (
	"GoTest/pkg/server"
	"log"
)

func main() {
	cnf := server.Config{
		Port: "8080",
	}

	svr, err := server.New(&cnf)
	switch err.(type) {
	case *server.InvalidConfig:
		log.Print("server setup failed") // NEED TO BE FIXED
	}
	svr.Start()
}
