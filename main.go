package bagit

import (
	"bufio"
	"crypto"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	// BagitVer is the version of the bagit spec this library coresponds to
	BagitVer = "0.97"
	// TagFileCharEnc is the encoding of the tag files
	TagFileCharEnc = "UTF-8"
)

// Bagit is a struct that describes a bag
type Bagit struct {
	Timestamp string
	Hashfunc  crypto.Hash
	Bagname   string
	Oxum      Oxum
}

// Oxum defnies a type that holds the sum of all bytes and files in the data dir
type Oxum struct {
	Bytes     int64
	Filecount int
}

// New creates a new Bagit struct
func New() *Bagit {
	return &Bagit{
		Timestamp: time.Now().Format("2006-01-02T15:04:05"),
	}
}

// Create creates a new bagit archive
func (b *Bagit) Create(srcDir string, outDir string, hashalg string) error {

	if hashalg == "md5" {
		log.Println("WARNING: md5 has known collisions. You should not use md5.")
		log.Println("WARNING: Press Ctrl + C to cancel or wait 5 seconds to continue...")
		time.Sleep(5 * time.Second)
	}

	// create bagit directory
	err := os.Mkdir(outDir, 0700)
	e(err)

	// create payload dir
	err = os.Mkdir(outDir+"/data", 0700)
	e(err)

	// create bagit.txt tag file
	fd, err := os.Create(outDir + "/bagit.txt")
	e(err)
	defer fd.Close()

	_, err = fd.WriteString("BagIt-Version: " + BagitVer + "\n")
	e(err)
	_, err = fd.WriteString("Tag-File-Character-Encoding: " + TagFileCharEnc)
	e(err)

	// create manifest-ALG.txt file
	fm, err := os.Create(outDir + "/manifest-" + hashalg + ".txt")
	e(err)
	defer fm.Close()

	// create bag-info.txt file
	fi, err := os.Create(outDir + "/bag-info.txt")
	e(err)
	defer fi.Close()

	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			b.Oxum.Filecount += 1
			fsize, err := os.Stat(path)
			e(err)
			b.Oxum.Bytes += fsize.Size()
			_, err = fm.WriteString(hex.EncodeToString(hashit(path, hashalg)) + " " + path + "\n")
			copy(path, outDir+"/data/"+path)

		} else {
			os.MkdirAll(outDir+"/data/"+path, 0700)
		}
		return nil
	})
	e(err)

	// write bag-info.txt
	oxumbytes := int(b.Oxum.Bytes)
	_, err = fi.WriteString("Bag-Software-Agent: bagit <https://github.com/steffenfritz/bagit>\n")
	_, err = fi.WriteString("Bagging-Date: " + b.Timestamp + "\n")
	_, err = fi.WriteString("Payload-Oxum: " + strconv.Itoa(oxumbytes) + "." + strconv.Itoa(b.Oxum.Filecount) + "\n")

	return nil
}

// Validate validates a bag for completeness and correctness
func (b *Bagit) Validate(srcDir string) error {
	var hashalg string
	var hashset bool

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(info.Name(), "manifest-") {
			if !hashset {
				hashalg = strings.Split(strings.Split(info.Name(), "-")[1], ".")[0]
				hashset = true
			}
		}
		return nil
	})
	e(err)

	// check oxum
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
			}
		}

	} else {
		log.Println("No bag-info.txt file found")
	}

	// get files from manifest file and calculate hash sum of files in bag and get info for oxum compare
	err = filepath.Walk(srcDir+"data/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			b.Oxum.Filecount += 1
			fsize, err := os.Stat(path)
			e(err)
			b.Oxum.Bytes += fsize.Size()
			hex.EncodeToString(hashit(path, hashalg))
		}
		return nil
	})
	e(err)

	oxumcalculated := strconv.Itoa(int(b.Oxum.Bytes)) + "." + strconv.Itoa(int(b.Oxum.Filecount))

	if oxumcalculated == oxumread {
		log.Println("Oxum valid")
	} else {
		log.Println("Oxum not valid")
	}

	return nil

}

// Tarit tars a directory
func (b *Bagit) Tarit() error {
	return nil
}
