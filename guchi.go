package main

import (
	"log"
	"os"
)

func main() {
	readConf()

	if err := cmd(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	os.Exit(0)
}
