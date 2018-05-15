package main

import (
	"log"

	"github.com/fenwickelliott/csync"
)

func main() {
	log.Fatal(csync.Serve(csync.Service{
		Name: "glam",
		Port: "4000",
		// Redirect: "http://localhost:5000",
		Redirect: "rock",
	}))
}
