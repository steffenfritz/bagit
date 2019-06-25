package bagit

import (
	"bufio"
	"encoding/hex"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Validate validates a bag for completeness and correctness
func (b *Bagit) Validate(srcDir string, verbose bool) error {
	var err error
	var hashalg string
	var hashset bool
	var manifestfile string
	var checkoxum bool
	bagvalid := true

	// filepath expects backslash
	if !strings.HasSuffix(srcDir, "/") {
		srcDir = srcDir + "/"
	}

	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(info.Name(), "manifest-") {
			if !hashset {
				hashalg = strings.Split(strings.Split(info.Name(), "-")[1], ".")[0]
				manifestfile = path
				hashset = true
			}
		}
		return err
	})
	e(err)

	if !hashset {
		log.Println("No manifest file found")
		log.Println("Bag not valid")
		return err
	}

	if verbose {
		log.Println("Used hash algorithm: " + hashalg)
	}

	// check oxum
	if verbose {
		log.Println("Looking for bag-info.txt file")
	}
	var oxumread string
	_, err = os.Stat(srcDir + "/bag-info.txt")
	if err == nil {
		fd, err := os.Open(srcDir + "/bag-info.txt")
		e(err)
		defer fd.Close()
		scanner := bufio.NewScanner(fd)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "Payload-Oxum:") {
				oxumread = strings.TrimSpace(strings.Split(scanner.Text(), ":")[1])
				checkoxum = true
			}
		}

	} else {
		log.Println("No bag-info.txt file found")

	}

	fm, err := os.Open(manifestfile)
	e(err)
	defer fm.Close()

	// walk through bag, calculate hashes and look up result in manifest file and get info for oxum compare
	if verbose {
		log.Println("Checking hashsums of files in payload directory")
	}

	err = filepath.Walk(srcDir+"data/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			b.Oxum.Filecount++
			fsize, err := os.Stat(path)
			e(err)
			b.Oxum.Bytes += fsize.Size()

			comppath := strings.SplitN(path, "/data/", 2)
			scanner := bufio.NewScanner(fm)
			fm.Seek(0, 0)
			var hashcorrect bool
			for scanner.Scan() {
				// normalizing strings here for comparison. We need a more elegant and faster way
				if strings.Join(strings.Fields(hex.EncodeToString(hashit(path, hashalg))+" "+comppath[1]), " ") == strings.Join(strings.Fields(scanner.Text()), " ") {
					hashcorrect = true
				}

			}
			if !hashcorrect {
				log.Println("File " + path + " not in manifest file or wrong hashsum!")
				bagvalid = false
			}

		}
		return nil
	})
	e(err)

	if checkoxum {
		oxumcalculated := strconv.Itoa(int(b.Oxum.Bytes)) + "." + strconv.Itoa(int(b.Oxum.Filecount))

		if oxumcalculated != oxumread {
			log.Println("Oxum not valid")
			bagvalid = false
		}

		if verbose {
			log.Println("Oxum in bag: \t" + oxumread)
			log.Println("Oxum calculated: \t" + oxumcalculated)
		}
	}
	if bagvalid == true {
		log.Println("Bag is valid.")
	} else {
		log.Println("Bag is not valid.")
	}

	return err

}

// ValidateFetchFile validates fetch.txt files for correct syntax
func ValidateFetchFile(inFetch string) bool {
	statFetchFile := true

	ff, err := os.Open(inFetch)
	e(err)
	scanner := bufio.NewScanner(ff)
	for scanner.Scan() {
		fetchuri := strings.Fields(scanner.Text())[0]
		fetchlen := strings.Fields(scanner.Text())[1]
		println(strings.Fields(scanner.Text())[2])

		// -- first field: check if uri format
		_, err := url.ParseRequestURI(fetchuri)
		if err != nil {
			log.Println("Fetch file contains at least one invalid URI. Quitting.")
			statFetchFile = false
			return statFetchFile
		}
		// -- second field: check if dash or number
		if fetchlen == "-" {
		}
		// -- third field: check if not empty
		//
	}
	return statFetchFile
}
