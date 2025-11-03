package main

import (
	"fmt"
	"log"

	"github.com/NKV510/url_saver/internal"
)

func main() {
	cfg, err := internal.Load()
	if err != nil {
		log.Panic("ERROR CONFIG: %w", err)
	}
	fmt.Println(cfg)
}
