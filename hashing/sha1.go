package hashing

import (
	"crypto/sha1"
	"math/big"
)

func ComputeSha1(s string) uint64 {
	h := sha1.New()
	h.Write([]byte(s))
	bi := big.NewInt(0)
	bi.SetBytes(h.Sum(nil))
	return bi.Uint64()
}
