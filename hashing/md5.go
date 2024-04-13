package hashing

import (
	"crypto/md5"
	"math/big"
)

func ComputeMD5(s string) uint64 {
	bi := big.NewInt(0)
	h := md5.New()
	h.Write([]byte(s))
	bi.SetBytes(h.Sum(nil))
	return bi.Uint64()
}
