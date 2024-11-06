package vdf

import (
	"math/big"

	"github.com/abraj/vdf/sqrt"
)

// Get a new VDF instance
func NewVDF() *sqrt.VDFSqrt {
	// p256 := "60464814417085833675395020742168312237934553084050601624605007846337253615407"

	// p257: 2**256 + 297
	p257 := "115792089237316195423570985008687907853269984665640564039457584007913129640233"

	p, _ := new(big.Int).SetString(p257, 10)
	return sqrt.NewVDFSqrt(p)
}
