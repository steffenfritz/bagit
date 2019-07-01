package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/steffenfritz/bagit"
)

const version = "0.2.0"

var starttime = time.Now().Format("2006-01-02T150405")

func main() {
	b := bagit.New()

	vers := flag.Bool("version", false, "Print version")
	validate := flag.String("validate", "", "Validate bag. Expects path to bag")
	b.SrcDir = flag.String("create", "", "Create bag. Expects path to source directory")
	b.OutDir = flag.String("output", "bag_"+starttime, "Output directory for bag. Used with create flag")
	tarit := flag.Bool("tar", false, "Create a tar archive when creating a bag")
	b.HashAlg = flag.String("hash", "sha512", "Hash algorithm used for manifest file when creating a bag [sha1, sha256, sha512, md5]")
	verbose := flag.Bool("v", false, "Verbose output")
	b.AddHeader = flag.String("header", "", "Additional headers for bag-info.txt. Expects path to json file")
	b.FetchFile = flag.String("fetch", "", "Adds optional fetch file to bag. Expects path to fetch.txt file and switch manifetch")
	b.FetchManifest = flag.String("manifetch", "", "Path to manifest file for optional fetch.txt file. Mandatory if fetch switch is used")

	flag.Parse()

	if *vers {
		log.Println("Version: " + version)

		return
	}

	if len(*validate) != 0 {
		b := bagit.New()
		_, err := os.Stat(*validate + "/fetch.txt")
		if err == nil {
			log.Println("Found a fetch.txt file in bag. Please add files before validating.")
		}
		bagvalid, err := b.Validate(*validate, *verbose)
		if err != nil {
			log.Println(err)
		}
		if !bagvalid {
			log.Println("Bag not valid.")
		} else {
			log.Println("Bag is valid.")
		}

		return
	}

	if len(*b.SrcDir) != 0 {
		_, err := os.Stat(*b.SrcDir)
		if err != nil {
			log.Println("Cannot read source directory")
			return
		}

		_, err = os.Stat(*b.OutDir)
		if err == nil {
			log.Println("Output directory already exists. Refusing to overwrite. Quitting.")
			return
		}
		// validate fetch.txt file and exit if not valid
		fetchStatus := false
		fetchoxumcompl := true
		fetchoxumbytes := 0
		fetchoxumfiles := 0

		if len(*b.FetchFile) != 0 {
			_, err := os.Stat(*b.FetchFile)
			if err != nil {
				log.Println("Could not read fetch.txt file. Quitting.")
				return
			}
			if len(*b.FetchManifest) == 0 {
				log.Println("The usage of a fetch.txt expects a related manifest file. Quitting.")
				return
			}

			fetchStatus, fetchoxumcompl, fetchoxumbytes, fetchoxumfiles = bagit.ValidateFetchFile(*b.FetchFile, *verbose)
			if !fetchStatus {
				log.Println("fetch.txt file not valid. Quitting.")
				return
			}

		}

		if !fetchoxumcompl {
			log.Println("fetch.txt: Using a dash in fetch.txt makes it impossible to update the oxum. Validation will fail.")
		}

		b.Oxum.Bytes = int64(fetchoxumbytes)
		b.Oxum.Filecount = fetchoxumfiles
		b.Create(*verbose)

		if *tarit {
			b.Tarit(*b.OutDir, *b.OutDir+".tar.gz")
		}

		return
	}

}
