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
func (b *Bagit) Validate(srcDir string, verbose bool) (bool, error) {

	var err error
	var hashalg string
	var hashset bool
	var manifestfile string
	var checkoxum bool
	var tagmanifests []string
	bagvalid := true

	// filepath expects backslash
	if !strings.HasSuffix(srcDir, string(os.PathSeparator)) {
		srcDir = srcDir + string(os.PathSeparator)
	}

	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(info.Name(), "manifest-") {
			if !hashset {
				hashalg = strings.Split(strings.Split(info.Name(), "-")[1], ".")[0]
				manifestfile = path
				hashset = true
			}
		}

		// add tagmanifests to list tagmanifests if present
		if strings.HasPrefix(info.Name(), "tagmanifest-") {
			tagmanifests = append(tagmanifests, info.Name())
		}
		return err
	})
	e(err)

	if !hashset {
		log.Println("No manifest file found")
		bagvalid = false
		return bagvalid, err
	}

	if verbose {
		log.Println("Used hash algorithm for payload manifest: " + hashalg)
	}

	// check oxum
	if verbose {
		log.Println("Looking for bag-info.txt file")
	}
	var oxumread string
	_, err = os.Stat(srcDir + string(os.PathSeparator) + "bag-info.txt")
	if err == nil {
		if verbose {
			log.Println("  Found bag-info.txt")
		}
		fd, err := os.Open(srcDir + string(os.PathSeparator) + "bag-info.txt")
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

	// checking if all files are present that are listed in payload manifest
	filescanner := bufio.NewScanner(fm)
	for filescanner.Scan() {
		tmpfile := strings.SplitN(filescanner.Text(), " ", 2)[1]
		_, err := os.Stat(srcDir + string(os.PathSeparator) + tmpfile)
		if err != nil {
			if verbose {
				log.Println(tmpfile + " is missing in the bag!")
			}
			bagvalid = false
		}
	}

	// walk through bag, calculate hashes and look up result in manifest file and get info for oxum compare
	if verbose {
		log.Println("Checking hashsums of files in payload directory")
	}

	err = filepath.Walk(srcDir+"data"+string(os.PathSeparator), func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			b.Oxum.Filecount++
			fsize, err := os.Stat(path)
			e(err)
			b.Oxum.Bytes += fsize.Size()

			comppath := strings.SplitN(path, string(os.PathSeparator)+"data"+string(os.PathSeparator), 2)
			scanner := bufio.NewScanner(fm)
			fm.Seek(0, 0)

			if verbose {
				log.Println("  Hashing " + path)
			}

			var hashcorrect bool
			for scanner.Scan() {
				calc := strings.Join(strings.Fields(hex.EncodeToString(hashit(path, hashalg))+" data"+string(os.PathSeparator)+comppath[1]), " ")
				read := strings.Join(strings.Fields(scanner.Text()), " ")
				if strings.EqualFold(calc, read) {
					hashcorrect = true
					return nil
				}
			}

			if !hashcorrect {
				if verbose {
					log.Println("File " + path + " not in manifest file or wrong hashsum!")
				}
				bagvalid = false
			}

		}
		return nil
	})
	e(err)

	if checkoxum {
		oxumcalculated := strconv.Itoa(int(b.Oxum.Bytes)) + "." + strconv.Itoa(int(b.Oxum.Filecount))

		if oxumcalculated != oxumread {
			if verbose {
				log.Println("Oxum not valid")
			}
			bagvalid = false
		}

		if verbose {
			log.Println("Oxum in bag: \t" + oxumread)
			log.Println("Oxum calculated: \t" + oxumcalculated)
		}
	}

	// validate tag manifests
	ValidateTagmanifests(&srcDir, &tagmanifests, verbose, &bagvalid)

	return bagvalid, err

}

// ValidateFetchFile validates fetch.txt files for correct syntax
// and returns sum of length and file count
func ValidateFetchFile(inFetch string, verbose bool) (bool, bool, int, int) {
	statFetchFile := true
	oxumlencomplete := true
	oxumbytes := 0
	oxumfiles := 0

	ff, err := os.Open(inFetch)
	e(err)
	scanner := bufio.NewScanner(ff)
	for scanner.Scan() {
		fetchuri := strings.Fields(scanner.Text())[0]
		fetchlen := strings.Fields(scanner.Text())[1]
		fetchpath := strings.Fields(scanner.Text())[2]

		// -- first field: check if uri format
		_, err := url.ParseRequestURI(fetchuri)
		if err != nil {
			if verbose {
				log.Println("fetch.txt: Fetch file contains at least one invalid URI. Quitting.")
			}
			statFetchFile = false
			return statFetchFile, false, 0, 0
		}
		// -- second field: check if dash or number
		if fetchlen != "-" {
			fileoxum, err := strconv.Atoi(fetchlen)
			if err != nil {
				if verbose {
					log.Println("fetch.txt: Length not a dash nor a number. Quitting.")
				}
				statFetchFile = false
				return statFetchFile, false, 0, 0
			}
			oxumbytes += fileoxum
		} else {
			oxumlencomplete = false
		}
		// -- third field: check if not empty
		if len(fetchpath) == 0 {
			if verbose {
				log.Println("fetch.txt: Local path empty. Quitting.")
			}
			statFetchFile = false
			return statFetchFile, false, 0, 0
		}
		oxumfiles++
	}
	return statFetchFile, oxumlencomplete, oxumbytes, oxumfiles
}

// ValidateTagmanifests validates the hash sums of tag files
func ValidateTagmanifests(srcDir *string, tagmanifests *[]string, verbose bool, bagvalid *bool) {
	if len(*tagmanifests) != 0 {
		for _, tmentry := range *tagmanifests {
			tmpfd, err := os.Open(*srcDir + string(os.PathSeparator) + tmentry)
			e(err)
			defer tmpfd.Close()
			tmphashalg := strings.Split(strings.Split(tmentry, "-")[1], ".")[0]

			if verbose {
				log.Println("Used hash algorithm for tagmanifest: " + tmphashalg)
			}

			scanner := bufio.NewScanner(tmpfd)
			for scanner.Scan() {
				tmptagfile := strings.Split(scanner.Text(), " ")[1]
				tmptagstat, err := os.Lstat(*srcDir + string(os.PathSeparator) + tmptagfile)
				e(err)
				if !tmptagstat.IsDir() {
					if verbose {
						log.Println("  Hashing " + *srcDir + tmptagfile)
					}
					calc := strings.Join(strings.Fields(hex.EncodeToString(hashit(*srcDir+tmptagfile, tmphashalg))+" "+tmptagfile), " ")
					read := strings.Join(strings.Fields(scanner.Text()), " ")
					if !strings.EqualFold(calc, read) {
						*bagvalid = false
						if verbose {
							log.Println("  File " + *srcDir + tmptagfile + " not in manifest file or wrong hashsum!")
						}
					}
				}
			}
		}
	}

}
