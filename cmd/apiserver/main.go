package main

import "log"

func main() {
	server := apiserver.New()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
