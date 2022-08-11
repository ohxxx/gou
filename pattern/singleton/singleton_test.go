package singleton

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	var s1, s2 *Singleton
	s1 = GetInstance()
	s2 = GetInstance()
	if s1 != s2 {
		t.Error("s1 != s2")
	}
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("GetInstance() != GetInstance()")
			}
		}
	})
}

func TestGetLazyInstance(t *testing.T) {
	var s1 *Singleton
	s1 = GetLazyInstance()
	s2 := GetLazyInstance()
	if s1 != s2 {
		t.Error("s1 != s2")
	}
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("GetLazyInstance() != GetLazyInstance()")
			}
		}
	})
}

func TestGetLazyInstanceParallel(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	const parCount = 100
	wg.Add(parCount)
	instances := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			//协程阻塞，等待channel被关闭才能继续运行
			<-start
			instances[index] = *GetLazyInstance()
			wg.Done()
		}(i)
	}
	//关闭channel，所有协程同时开始运行，实现并行(parallel)
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
