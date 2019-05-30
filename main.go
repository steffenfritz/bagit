package bagit

import (
	"crypto"
	"log"
	"os"
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
	Oxum      string
}

// New creates a new Bagit struct
func New() *Bagit {
	return &Bagit{
		Timestamp: time.Now().Format("2006-01-02T15:04:05"),
	}
}

// Create creates a new bagit archive
func (b *Bagit) Create(srcDir string, outDir string, hashalg string) error {
	// create bagit directory
	err := os.Mkdir(outDir, 0700)
	e(err)

	// create payload dir
	err = os.Mkdir(outDir+"/data", 0700)
	e(err)

	// create bagit.txt tag file
	fd, err := os.Create(outDir + "/bagit.txt")
	e(err)

	_, err = fd.WriteString("BagIt-Version: " + BagitVer + "\n")
	e(err)
	_, err = fd.WriteString("Tag-File-Character-Encoding: " + TagFileCharEnc)
	e(err)

	return nil
}

// Validate validates a bag for completeness and correctness
func (b *Bagit) Validate() error {
	return nil
}

// Tarit tars a directory
func (b *Bagit) Tarit() error {
	return nil
}

// e is just a shorty for generic errors and panics
func e(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
