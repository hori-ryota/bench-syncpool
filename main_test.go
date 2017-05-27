package main

import (
	"strconv"
	"sync"
	"testing"
)

type Chunk struct {
	SeqId    uint64
	URI      string
	Duration float64
}

func BenchmarkWithoutPool(b *testing.B) {
	testCounts := []int{0, 1, 10, 100, 1000, 10000}
	for _, count := range testCounts {
		name := strconv.Itoa(count)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = make([]Chunk, count)
			}
		})
	}
}

func BenchmarkWithPool(b *testing.B) {
	testCounts := []int{0, 1, 10, 100, 1000, 10000}
	for _, count := range testCounts {
		name := strconv.Itoa(count)
		b.Run(name, func(b *testing.B) {
			pool := sync.Pool{
				New: func() interface{} {
					return make([]Chunk, count)
				},
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				c := pool.Get().([]Chunk)
				pool.Put(c)
			}
		})
	}
}
