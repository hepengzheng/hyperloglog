package gohll

import (
	"context"
	"hash/fnv"
	"math"
	"math/bits"

	"github.com/hepengzheng/gohll/hashing"
)

const (
	b                       = 14
	m                       = 1 << b
	alpha                   = 0.7213 / (1 + 1.079/float64(m))
	indexBitPattern         = m - 1
	upperBoundRelativeError = 1.04 / (1 << (b >> 1))
)

type HyperLogLog interface {
	Add(ctx context.Context, s string)
	Count(ctx context.Context) int
}

var _ HyperLogLog = (*MyHLL)(nil)

type MyHLL struct {
	buckets [m]int

	hashAlg hashing.HashAlg
}

func NewHyperLogLog(hashAlg hashing.HashAlg) *MyHLL {
	return &MyHLL{
		hashAlg: hashAlg,
	}
}

func (hll *MyHLL) Add(_ context.Context, value string) {
	x := hll.hashAlg(value)
	index := getIndex(x)
	pho := computeRho(x)
	if pho > hll.buckets[index] {
		hll.buckets[index] = pho
	}
}

func (hll *MyHLL) Count(_ context.Context) int {
	e := hll.rawEstimate()
	if e <= float64(m)*2.5 {
		if v := hll.linearCounting(); v != 0 {
			e = float64(m) * math.Log2(float64(m)/float64(v))
		}
	}
	return int(e)
}

func (hll *MyHLL) linearCounting() int {
	v := 0
	for _, bucket := range hll.buckets {
		if bucket == 0 {
			v++
		}
	}
	return v
}

func (hll *MyHLL) rawEstimate() float64 {
	var sum float64
	for _, bucket := range hll.buckets {
		sum += math.Pow(2, -float64(bucket))
	}
	return alpha * float64(m*m) / sum
}

func computeRho(x uint64) int {
	return 1 + bits.TrailingZeros64(x>>b)
}

func getIndex(x uint64) int {
	return int(x & indexBitPattern)
}
