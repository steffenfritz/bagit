package bagit

import "testing"

func TestValidateFetchFile(t *testing.T) {
	table := []struct {
		inFile       string
		fetchValid   bool
		oxumComplete bool
		oxumBytes    int
		oxumFiles    int
	}{
		{"testdata/fetch.txt", true, true, 2783, 2},
		{"testdata/fetch_dash.txt", true, false, 2783, 3},
		{"testdata/fetch_fail.txt", false, false, 0, 0},
	}

	for _, entry := range table {
		fetchValid, oxumComplete, oxumBytes, oxumFiles := ValidateFetchFile(entry.inFile)
		if entry.fetchValid != fetchValid {
			t.Errorf(entry.inFile+" = %d", fetchValid)
		}
		if entry.oxumComplete != oxumComplete {
			t.Errorf(entry.inFile+" = %d", oxumComplete)
		}
		if entry.oxumBytes != oxumBytes {
			t.Errorf(entry.inFile+" = %d", oxumBytes)
		}
		if entry.oxumFiles != oxumFiles {
			t.Errorf(entry.inFile+" = %d", oxumFiles)
		}
	}
}
