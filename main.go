package main

import (
	"GoTest/pkg/server"
	"fmt"
)

func main() {
	cnf := server.Config{
		Port: "8080",
	}

	s, err := server.New(&cnf)
	switch err.(type) {
	case *server.InvalidConfig:
		fmt.Errorf("server setup failed") // NEED TO BE FIXED
	}
	s.Start()
}
