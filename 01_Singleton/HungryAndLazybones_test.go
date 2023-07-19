package _1_Singleton

import (
	"testing"
)

func TestGetHungryLazybonesInstance(t *testing.T) {
	ins1 := GetHungryLazybonesInstance()
	ins2 := GetHungryLazybonesInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func BenchmarkGetHungryLazybonesInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetHungryLazybonesInstance() != GetHungryLazybonesInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
