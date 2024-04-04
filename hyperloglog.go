package gohll

import (
	"context"
	"hash/fnv"
	"math"
	"math/bits"
)

const (
	b                       = 14
	m                       = 1 << b
	alpha                   = 0.7213 / (1 + 1.079/float64(m))
	upperBoundRelativeError = 1.04 / (1 << (b >> 1))
)

type HyperLogLog interface {
	Add(ctx context.Context, s string)
	Count(ctx context.Context) int
}

var _ HyperLogLog = (*MyHLL)(nil)

type MyHLL struct {
	buckets []int
}

func NewHyperLogLog() *MyHLL {
	return &MyHLL{
		buckets: make([]int, m),
	}
}

func (hll *MyHLL) Add(_ context.Context, value string) {
	x := computeHash(value)
	index := hll.getIndex(x)
	pho := computePho(x >> b)
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
	} else if e >= (1<<32)/30 {
		e = -(1 << 32) * math.Log2(1-e/float64(1<<32))
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
	estimate := alpha * float64(m*m) / sum
	return estimate
}

func computeHash(s string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

func computePho(x uint64) int {
	return 1 + bits.TrailingZeros64(x)
}

func (hll *MyHLL) getIndex(x uint64) int {
	return int(x & 0x3FFF)
}
