package bagit

import (
	"os"
	"testing"
)

func TestGetaddHeader(t *testing.T) {
	result := getaddHeader("testdata/header.json")

	v := result["Source-Organization"]
	if v.(string) != "ACME GmbH AG Inc" {
		t.Errorf("%s", v.(string))
	}

	v = result["Organization-Address"]
	if v.(string) != "digipress@fritz.wtf" {
		t.Errorf("%s", v.(string))
	}

}

func TestCreate(t *testing.T) {
	// command line flags
	var verbose bool
	srcDir := "testdata/testinput"
	psrcDir := &srcDir
	outDir := "testdata/bag_golden"
	poutDir := &outDir
	addHeader := "testdata/header.json"
	paddHeader := &addHeader
	hashAlg := "sha512"
	phashAlg := &hashAlg
	fetchFile := ""
	pfetchFile := &fetchFile
	fetchManifest := ""
	pfetchManifest := &fetchManifest

	b := New()
	b.SrcDir = psrcDir
	b.OutDir = poutDir
	b.AddHeader = paddHeader
	b.HashAlg = phashAlg
	b.FetchFile = pfetchFile
	b.FetchManifest = pfetchManifest
	b.Create(verbose)

	_, err := os.Open("testdata/bag_golden/data/testdata/testinput/random.data")
	if err != nil {
		t.Errorf("Payload file was not copied")
	}

	_, err = os.Open("testdata/bag_golden/bag-info.txt")
	if err != nil {
		t.Errorf("bag-info.txt was not created")
	}

	_, err = os.Open("testdata/bag_golden/bagit.txt")
	if err != nil {
		t.Errorf("bagit.txt was not created")
	}

	_, err = os.Open("testdata/bag_golden/manifest-sha512.txt")
	if err != nil {
		t.Errorf("manifest-sha512.txt was not created")
	}

	// reset oxum for validation
	b.Oxum.Bytes = 0
	b.Oxum.Filecount = 0

	valid, err := b.Validate(outDir, false)
	if !valid {
		t.Errorf("Created bag not valid")
	}
	if err != nil {
		t.Error(err)
	}
}
