package main

import (
	"flag"
	"log"
	"time"

	"github.com/steffenfritz/bagit"
)

const version = "0.1.0-DEV"

var starttime = time.Now().Format("2006-01-02T150405")

func main() {

	vers := flag.Bool("version", false, "Print version")
	validate := flag.String("validate", "", "Validate bag. Expects path to bag")
	createSrc := flag.String("create", "", "Create bag. Expects path to source directory")
	outputDir := flag.String("output", "bag_"+starttime, "Output directory for bag. Used with create flag")
	tarit := flag.Bool("tar", false, "Create a tar archive when creating a bag")
	hashalg := flag.String("hash", "sha256", "Hash algorithm used for manifest file when creating a bag [sha1, sha256, md5]")
	verbose := flag.Bool("v", false, "Verbose output")
	addHeader := flag.String("H", "", "Additional headers for bag-info.txt. Expects path to json file")

	flag.Parse()

	if *vers {
		log.Println("Version: " + version)

		return
	}

	if len(*validate) != 0 {
		b := bagit.New()
		b.Validate(*validate, *verbose)

		return
	}

	if len(*createSrc) != 0 {
		b := bagit.New()
		b.Create(*createSrc, *outputDir, *hashalg, *addHeader)

		if *tarit {
			b.Tarit(*outputDir, *outputDir+".tar.gz")
		}

		return
	}

}
