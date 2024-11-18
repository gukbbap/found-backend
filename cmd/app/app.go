package main

import (
	"fmt"
	"found-backend/internal/infra/config"
	"log"
)

func main() {
	found, err := config.NewConfig(".env.toml")
	if err != nil {
		log.Panicf("failed to new config: %v\n", err)
	}

	fmt.Println("config struct: ", found)
}
