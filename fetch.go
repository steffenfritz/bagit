package bagit

import (
	"log"
	"time"
)

func fetchCreate(fetchFile string, fetchManifest *map[string]string) {

	// Download src file
	// -- http
	// -- ftp
	// -- ssh
	// -- smb
	// else : not implemented

	// calculate hashsum

	// return map hashsum filename for writing
}

func fetchValidate(fetchFile string) {

	log.Println("Validating a bag with a fetch.txt file completes your bag.\nThis may take a while. Cancel with Ctrl+C or wait 3 seconds")
	time.Sleep(3 * time.Second)
	// NEXT
	// read lines, returns int and map srcurl and filename
	log.Println("Downloading %i files")
	// Download src file
	// -- http
	// -- ftp
	// -- ssh
	// -- smb
	// else : not implemented

	// write to destination bag

}
