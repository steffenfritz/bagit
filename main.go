package bagit

import (
	"crypto"
	"time"
)

const (
	// BagitVer is the version of the bagit spec this library coresponds to
	BagitVer = "0.97"
	// TagFileCharEnc is the encoding of the tag files
	TagFileCharEnc = "UTF-8"
)

// Bagit is a struct that describes a bag
type Bagit struct {
	Timestamp string
	Hashfunc  crypto.Hash
	Bagname   string
	Oxum      string
}

// New creates a new Bagit struct
func New() *Bagit {
	return &Bagit{
		Timestamp: time.Now().Format("2006-01-02T15:04:05"),
	}
}

// Create creates a new bagit archive
func (b *Bagit) Create(srcDir string, outDir string, hashalg string) error {
	return nil
}

// Validate validates a bag for completeness and correctness
func (b *Bagit) Validate() error {
	return nil
}

// Tarit tars a directory
func (b *Bagit) Tarit() error {
	return nil
}
