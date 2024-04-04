package gohll

import (
	"bufio"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHyperLogLog(t *testing.T) {
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
	myhll := NewHyperLogLog()
	redishll := NewRedisHLL()
	for scanner.Scan() {
		line := scanner.Text()
		myhll.Add(ctx, line)
		redishll.Add(ctx, line)
	}

	count := myhll.Count(ctx)
	fmt.Printf("count: %d\n", count)
	relativeErr := relativeError(count)

	redisCount := redishll.Count(ctx)
	fmt.Printf("redis count: %d\n", redisCount)
	redisRelativeErr := relativeError(redisCount)

	t.Cleanup(func() {
		redishll.rdb.Del(ctx, redisKey)
	})

	assert.True(t, relativeErr < upperBoundRelativeError)

	fmt.Printf("relative error of my implementation: %f\n", relativeErr)
	fmt.Printf("relative error of Redis' implementation: %f\n", redisRelativeErr)
}
