package main

import (
	"bomberman-dom/internal/app/apiserver"
	"log"
)

func main() {
	cfg := apiserver.LoadConfig()
	if err := apiserver.Start(cfg); err != nil {
		log.Fatal(err)
	}

}
