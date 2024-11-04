package vdf

import (
	"math/big"

	"github.com/abraj/vdf/sqrt"
)

// Get a new VDF instance
func NewVDF() *sqrt.VDFSqrt {
	p256 := "60464814417085833675395020742168312237934553084050601624605007846337253615407"
	p, _ := new(big.Int).SetString(p256, 10)
	return sqrt.NewVDFSqrt(p)
}
