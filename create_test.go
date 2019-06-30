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
