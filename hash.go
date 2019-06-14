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
	e(err)
	defer fd.Close()

	var hasher hash.Hash

	if hashalg == "sha256" {
		hasher = sha256.New()
		_, err = io.Copy(hasher, fd)
		e(err)

		checksum := hasher.Sum(nil)

		return checksum

	} else if hashalg == "md5" {
		hasher = md5.New()
		_, err = io.Copy(hasher, fd)
		e(err)

		checksum := hasher.Sum(nil)

		return checksum

	} else if hashalg == "sha1" {
		hasher = sha1.New()
		_, err = io.Copy(hasher, fd)
		e(err)

		checksum := hasher.Sum(nil)

		return checksum

	} else if hashalg == "sha512" {
		hasher = sha512.New()
		_, err = io.Copy(hasher, fd)
		e(err)

		checksum := hasher.Sum(nil)

		return checksum

	} else {
		log.Println("Hash not implemented")
		os.Exit(1)
	}

	return nil
}
