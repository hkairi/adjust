package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hkairi/adjust/fetcher"
)

func main() {
	limit := flag.Int("limit", 10, "Limit of parallel tasks.")
	flag.Parse()

	liste := os.Args[3:]

	if *limit > 10 {
		fmt.Println("Limit too high")
		os.Exit(0)
	}

	adjust := fetcher.New(*limit, liste)
	adjust.Start()
}
