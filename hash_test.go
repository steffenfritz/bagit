package bagit

import (
	"bytes"
	"testing"
)

func TestHashit(t *testing.T) {

	table := []struct {
		algo      string
		shouldsum []byte
	}{
		{"md5", []byte{216, 124, 3, 107, 198, 190, 85, 235, 207, 28, 124, 96, 252, 96, 138, 223}},
		{"sha1", []byte{240, 70, 94, 175, 229, 216, 83, 113, 218, 150, 50, 145, 44, 191, 241, 188, 57, 168, 220, 165}},
		{"sha256", []byte{46, 183, 100, 155, 96, 186, 38, 2, 166, 69, 219, 252, 20, 54, 49, 85, 46, 105, 103, 210, 42, 65, 170, 208, 29, 165, 134, 123, 121, 111, 193, 21}},
		{"sha512", []byte{112, 15, 179, 86, 7, 165, 50, 179, 183, 215, 116, 86, 2, 207, 152, 92, 97, 207, 18, 181, 172, 241, 27, 181, 85, 132, 32, 67, 99, 157, 9, 251, 98, 62, 27, 143, 86, 193, 13, 95, 116, 142, 179, 194, 180, 118, 71, 120, 223, 58, 13, 9, 28, 36, 77, 232, 70, 145, 15, 95, 72, 37, 168, 57}},
	}

	for _, entry := range table {

		checksum := hashit("testdata/random.data", entry.algo)
		bytes.Equal(checksum, entry.shouldsum)
		if bytes.Equal(checksum, entry.shouldsum) == false {
			t.Errorf("testdata/random.data = %d", checksum)
		}
	}

}
