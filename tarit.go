package bagit

import (
	"github.com/mholt/archiver"
)

// Tarit tars a directory
func (b *Bagit) Tarit(srcDir string, outFile string) error {

	tarbag(srcDir, outFile)
	return nil
}

func tarbag(src string, outarc string) error {

	err := archiver.Archive([]string{src}, outarc)
	e(err)
	return err
}
