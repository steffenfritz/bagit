package bagit

import (
	"fmt"
	"io"
	"os"
)

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	e(err)

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	e(err)
	defer source.Close()

	destination, err := os.Create(dst)
	e(err)
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
