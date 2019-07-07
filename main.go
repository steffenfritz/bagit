package bagit

import (
	"crypto"
	"time"
)

const (
	// BagitVer is the version of the bagit spec this library coresponds to
	BagitVer = "1.0"
	// TagFileCharEnc is the encoding of the tag files
	TagFileCharEnc = "UTF-8"
)

// Bagit is a struct that describes a bag
type Bagit struct {
	Timestamp     string
	Hashfunc      crypto.Hash
	Bagname       string
	Oxum          Oxum
	SrcDir        *string
	OutDir        *string
	HashAlg       *string
	FetchFile     *string
	FetchManifest *string
	AddHeader     *string
	TagManifest   *string
}

// Oxum defnies a type that holds the sum of all bytes and files in the data dir
type Oxum struct {
	Bytes     int64
	Filecount int
}

// New creates a new Bagit struct
func New() *Bagit {
	return &Bagit{
		Timestamp: time.Now().Format("2006-01-02T15:04:05"),
	}
}

// Tarit tars a directory
func (b *Bagit) Tarit(srcDir string, outFile string) error {

	tarbag(srcDir, outFile)
	return nil
}
