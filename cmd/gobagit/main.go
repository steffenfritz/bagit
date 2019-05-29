package main

import (
	"flag"
	"log"

	"github.com/steffenfritz/bagit"
)

const version = "0.1.0-DEV"

func main() {

	vers := flag.Bool("version", false, "Print version")
	validate := flag.String("validate", "", "Validate bag. Expects path to bag")
	createSrc := flag.String("create", "", "Create bag. Expects path to source directory")
	outputDir := flag.String("output", "", "Output directory for bag. Used with create flag")
	tarit := flag.Bool("tar", false, "Create a tar archive when creating a bag")
	hashalg := flag.String("hash", "sha256", "Hash algorithm used for manifest file. Defaults to SHA256")

	flag.Parse()

	if *vers {
		log.Println("Version: " + version)
		return
	}

	if len(*validate) != 0 {
		b := bagit.New()
		b.Validate()

		return
	}

	if len(*createSrc) != 0 {
		b := bagit.New()
		b.Create(*createSrc, *outputDir, *hashalg)

		if *tarit {
			b.Tarit()
		}

		return
	}

}
