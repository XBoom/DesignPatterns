package _1_Singleton

import (
	"testing"
)

func TestGetHungryInstance(t *testing.T) {
	ins1 := GetHungryInstance()
	ins2 := GetHungryInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func BenchmarkGetHungryInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetHungryInstance() != GetHungryInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
