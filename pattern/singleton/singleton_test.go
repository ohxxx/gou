package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	var s1, s2 *singleton.Singleton
	s1 = singleton.GetInstance()
	s2 = singleton.GetInstance()
	if s1 != s2 {
		t.Error("s1 != s2")
	}
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetInstance() != singleton.GetInstance() {
				b.Errorf("singleton.GetInstance() != singleton.GetInstance()")
			}
		}
	})
}

func TestGetLazyInstance(t *testing.T) {
	var s1, s2 *singleton.Singleton
	s1 = singleton.GetLazyInstance()
	s2 = singleton.GetLazyInstance()
	if s1 != s2 {
		t.Error("s1 != s2")
	}
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetLazyInstance() != singleton.GetLazyInstance() {
				b.Errorf("singleton.GetLazyInstance() != singleton.GetLazyInstance()")
			}
		}
	})
}
