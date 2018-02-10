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
				_ = make([]Chunk, 0, count)
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
					return make([]Chunk, 0, count)
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

// func BenchmarkWithPoolAndInitStruct(b *testing.B) {
// 	testCounts := []int{0, 1, 10, 100, 1000, 10000}
// 	for _, count := range testCounts {
// 		name := strconv.Itoa(count)
// 		b.Run(name, func(b *testing.B) {
// 			pool := sync.Pool{
// 				New: func() interface{} {
// 					return make([]Chunk, count)
// 				},
// 			}
// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				c := pool.Get().([]Chunk)
// 				for j := range c {
// 					c[j] = Chunk{}
// 				}
// 				pool.Put(c)
// 			}
// 		})
// 	}
// }

// func BenchmarkWithPoolAndInitFields(b *testing.B) {
// 	testCounts := []int{0, 1, 10, 100, 1000, 10000}
// 	for _, count := range testCounts {
// 		name := strconv.Itoa(count)
// 		b.Run(name, func(b *testing.B) {
// 			pool := sync.Pool{
// 				New: func() interface{} {
// 					return make([]Chunk, count)
// 				},
// 			}
// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				c := pool.Get().([]Chunk)
// 				for j := range c {
// 					c[j].SeqId = 0
// 					c[j].Duration = 0
// 					c[j].URI = ""
// 				}
// 				pool.Put(c)
// 			}
// 		})
// 	}
// }

// func (c *Chunk) Init() {
// 	c.SeqId = 0
// 	c.Duration = 0
// 	c.URI = ""
// }

// func BenchmarkWithPoolAndInitFunc(b *testing.B) {
// 	testCounts := []int{0, 1, 10, 100, 1000, 10000}
// 	for _, count := range testCounts {
// 		name := strconv.Itoa(count)
// 		b.Run(name, func(b *testing.B) {
// 			pool := sync.Pool{
// 				New: func() interface{} {
// 					return make([]Chunk, count)
// 				},
// 			}
// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				c := pool.Get().([]Chunk)
// 				for j := range c {
// 					c[j].Init()
// 				}
// 				pool.Put(c)
// 			}
// 		})
// 	}
// }

// func BenchmarkWithPoolAndSetLenZero(b *testing.B) {
// 	testCounts := []int{0, 1, 10, 100, 1000, 10000}
// 	for _, count := range testCounts {
// 		name := strconv.Itoa(count)
// 		b.Run(name, func(b *testing.B) {
// 			pool := sync.Pool{
// 				New: func() interface{} {
// 					return make([]Chunk, count)
// 				},
// 			}
// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				c := pool.Get().([]Chunk)
// 				c = c[0:0]
// 				pool.Put(c)
// 			}
// 		})
// 	}
// }
