package bagit

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"io"
	"log"
	"os"
)

func hashit(inFile string, hashalg string) []byte {
	fd, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {
			log.Printf("WARNING: %s", err.Error())
		}
	}(fd)

	var hasher hash.Hash

	if hashalg == "sha256" {
		hasher = sha256.New()

	} else if hashalg == "md5" {
		hasher = md5.New()

	} else if hashalg == "sha1" {
		hasher = sha1.New()

	} else if hashalg == "sha512" {
		hasher = sha512.New()

	} else {
		log.Println("Hash not implemented")
		os.Exit(1)
	}

	_, err = io.Copy(hasher, fd)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	checksum := hasher.Sum(nil)

	return checksum

}
