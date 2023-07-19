package _1_Singleton

import (
	"testing"
)

func TestGetLazybonesInstance(t *testing.T) {
	ins1 := GetLazybonesInstance()
	ins2 := GetLazybonesInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func BenchmarkGetLazybonesInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazybonesInstance() != GetLazybonesInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
