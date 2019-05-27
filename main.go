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
	Timestamp time.Time
	Hashfunc  crypto.Hash
	Bagname   string
	Oxum      string
}

// New creates a new Bagit struct
func New() *Bagit {
	return &Bagit{
		Timestamp: time.Now(),
	}
}

// Create creates a new bagit archive
func (b *Bagit) Create() error {
	return nil
}

// Validate validates a bag for completeness and correctness
func (b *Bagit) Validate() error {
	return nil
}
