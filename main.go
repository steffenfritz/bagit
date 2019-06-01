package bagit

import (
	"crypto"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
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

	if hashalg == "md5" {
		log.Println("WARNING: md5 has known collisions. You should not use md5.")
		log.Println("WARNING: Press Ctrl + C to cancel or wait 5 seconds to continue...")
		time.Sleep(5 * time.Second)
	}
	//filehashmap := make(map[string]string)

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

	fe, err := os.Create(outDir + "/manifest-" + hashalg + ".txt")
	e(err)
	defer fe.Close()

	_, err = fd.WriteString("BagIt-Version: " + BagitVer + "\n")
	e(err)
	_, err = fd.WriteString("Tag-File-Character-Encoding: " + TagFileCharEnc)
	e(err)

	fm, err := os.Create(outDir + "/manifest-" + hashalg + ".txt")
	e(err)
	defer fm.Close()

	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			_, err = fm.WriteString(hex.EncodeToString(hashit(path, hashalg)) + " " + path + "\n")
			// NEXT

			copy(path, outDir+"/data"+path)
		} else {
			os.MkdirAll(outDir+"/data"+path, 0700)
		}
		return nil
	})
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
