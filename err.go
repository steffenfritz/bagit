package bagit

import "log"

// e is just a shorty for generic errors and panics
func e(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
