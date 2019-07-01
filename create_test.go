package bagit

import "testing"

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
	srcDir := "testdata/Bag_Input"
	psrcDir := &srcDir
	outDir := "testdata/Bag_golden"
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
}
