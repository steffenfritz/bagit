package bagit

import (
	"crypto"
)

//goland:noinspection GoNameStartsWithPackageName
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
