package bagit

import (
	"fmt"
	"io"
	"log"
	"os"
)

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			log.Printf("WARNING: %s", err.Error())
		}
	}(source)

	destination, err := os.Create(dst)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	defer func(destination *os.File) {
		err := destination.Close()
		if err != nil {
			log.Printf("WARNING: %s", err.Error())
		}
	}(destination)

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
