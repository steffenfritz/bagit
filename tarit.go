package bagit

import (
	"github.com/mholt/archiver"
	"log"
)

// Tarit tars a directory
func (b *Bagit) Tarit(srcDir string, outFile string) error {
	err := tarbag(srcDir, outFile)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	return nil
}

func tarbag(src string, outarc string) error {
	err := archiver.Archive([]string{src}, outarc)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	return err
}
