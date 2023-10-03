package main

import "github.com/dstan05/auth/internal/server"

func main() {
	s, err := server.Init()
	if err != nil {
		panic(err)
	}
	defer s.Stop()
}
