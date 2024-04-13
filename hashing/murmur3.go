package hashing

import (
	"math/big"

	"github.com/spaolacci/murmur3"
)

func ComputeMurmur3(s string) uint64 {
	h := murmur3.New64()
	h.Write([]byte(s))
	bi := big.NewInt(0)
	bi.SetBytes(h.Sum(nil))
	return bi.Uint64()
}
