package main

import (
	"flag"

	_ "github.com/steffenfritz/bagit"
)

func main() {

	sourcePath := flag.String("source", "", "Source directory for creating a bag")

	flag.Parse()
}
