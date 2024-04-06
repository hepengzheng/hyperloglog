package gohll

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	myhll    = NewHyperLogLog()
	redishll = NewRedisHLL()
)

func init() {
	// Uncomment the following line to regenerate the input set
	// and this is necessary when you run the test for the first time
	// note that the generating process may be time-consuming,
	// you may want to reduce the processing time by adjusting
	// the value of 'numStrings' at the top of the file.

	//Init()

	file, _ := os.Open(filename)
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	ctx := context.Background()
	redishll.rdb.Del(ctx, redisKey)
	for scanner.Scan() {
		line := scanner.Text()
		myhll.Add(ctx, line)
		redishll.Add(ctx, line)
	}

}

func Test_MyHLL(t *testing.T) {
	ctx := context.Background()
	count := myhll.Count(ctx)
	fmt.Printf("count: %d\n", count)
	relativeErr := relativeError(count)
	assert.True(t, relativeErr < upperBoundRelativeError)
	fmt.Printf("relative error of my implementation: %f\n", relativeErr)
}

func Benchmark_MyHLL(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		_ = myhll.Count(ctx)
	}
}

func Test_RedisHLL(t *testing.T) {
	ctx := context.Background()
	redisCount := redishll.Count(ctx)
	fmt.Printf("redis count: %d\n", redisCount)
	redisRelativeErr := relativeError(redisCount)
	fmt.Printf("relative error of Redis' implementation: %f\n", redisRelativeErr)
}

func Benchmark_RedisHLL(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		_ = redishll.Count(ctx)
	}
}
