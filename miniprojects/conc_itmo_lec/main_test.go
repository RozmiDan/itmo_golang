package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomicSpeed(b *testing.B) {
	var fstVal int
	var scndVal atomic.Int64

	b.Run("with mutex", func(b *testing.B) {
		mx := &sync.Mutex{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			mx.Lock()
			fstVal++
			mx.Unlock()
		}
	})

	b.Run("with atomic", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			scndVal.Add(1)
		}
	})
}