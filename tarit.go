package bagit

import (
	"github.com/mholt/archiver"
)

func tarbag(src string, outarc string) error {

	err := archiver.Archive([]string{src}, outarc)
	e(err)
	return err
}
