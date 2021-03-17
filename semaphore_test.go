package semaphore

import (
	"sync/atomic"
	"testing"
)

func BenchmarkSemaphoreAcquire1(b *testing.B) {
	sema := NewSemaphore(1)
	var count int64 = 0
	for i:=0; i< b.N; i++{
		go func() {
			sema.Acquire()
			defer func() {
				sema.Release()
			}()
			atomic.AddInt64(&count,1)
		}()
	}
}

func BenchmarkSemaphoreAcquire2(b *testing.B) {
	sema := NewSemaphore(2)
	var count int64 = 0
	for i:=0; i< b.N; i++{
		go func() {
			sema.Acquire()
			defer func() {
				sema.Release()
			}()
			atomic.AddInt64(&count,1)
		}()
	}
}

func BenchmarkSemaphoreAcquire3(b *testing.B) {
	sema := NewSemaphore(3)
	var count int64 = 0
	for i:=0; i< b.N; i++{
		go func() {
			sema.Acquire()
			defer func() {
				sema.Release()
			}()
			atomic.AddInt64(&count,1)
		}()
	}
}

func BenchmarkSemaphoreAcquire4(b *testing.B) {
	sema := NewSemaphore(4)
	var count int64 = 0
	for i:=0; i< b.N; i++{
		go func() {
			sema.Acquire()
			defer func() {
				sema.Release()
			}()
			atomic.AddInt64(&count,1)
		}()
	}
}

func BenchmarkAdd(b *testing.B) {
	var count int64 = 0
	for i:=0; i< b.N; i++{
		go func() {
			atomic.AddInt64(&count,1)
		}()
	}
}