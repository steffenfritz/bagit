package bagit

import "testing"

func TestValidateFetchFile(t *testing.T) {
	table := []struct {
		inFile       string
		verbose      bool
		fetchValid   bool
		oxumComplete bool
		oxumBytes    int
		oxumFiles    int
	}{
		{"testdata/fetch.txt", false, true, true, 2783, 2},
		{"testdata/fetch_dash.txt", false, true, false, 2783, 3},
		{"testdata/fetch_fail.txt", false, false, false, 0, 0},
		{"testdata/fetch_nan.txt", false, false, false, 0, 0},
	}

	for _, entry := range table {
		fetchValid, oxumComplete, oxumBytes, oxumFiles := ValidateFetchFile(entry.inFile, entry.verbose)
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

func TestValidate(t *testing.T) {
	b := New()
	table := []struct {
		inBag     string
		verbose   bool
		bagValid  bool
		someError error
	}{
		{"testdata/LoC_Bag_01", false, true, nil},
		{"testdata/LoC_Bag_02", false, false, nil},
	}

	for _, entry := range table {
		bagValid, someError := b.Validate(entry.inBag, entry.verbose)
		if entry.bagValid != bagValid {
			t.Errorf(entry.inBag+" = %d", bagValid)
		}
		if entry.someError != someError {
			t.Error(someError)
		}
	}
}
