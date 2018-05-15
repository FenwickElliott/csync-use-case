package main

import (
	"log"

	"github.com/fenwickelliott/csync"
)

func main() {
	log.Fatal(csync.Serve(csync.Service{
		Name:     "rock",
		Port:     "5000",
		Redirect: "http://localhost:4000",
	}))
}
