package hashing

import (
	"crypto/sha256"
	"math/big"
)

func ComputeSha256(s string) uint64 {
	h := sha256.New()
	h.Write([]byte(s))
	bi := big.NewInt(0)
	bi.SetBytes(h.Sum(nil))
	return bi.Uint64()
}
