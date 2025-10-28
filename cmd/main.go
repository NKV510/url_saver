package main

import (
	"fmt"

	"github.com/NKV510/url_saver/internal"
)

func main() {
	cfg := internal.MustLoad()
	fmt.Println(cfg)
}
