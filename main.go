package main

import "log"

func main() {
	if err := StartServer(); err != nil {
		log.Fatalf("Server failed to start %s", err)
	}
}