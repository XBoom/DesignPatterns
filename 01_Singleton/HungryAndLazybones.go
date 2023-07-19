package _1_Singleton

import "sync"

// HungryLazybones 懒汉模式与饿汉模式结合
type HungryLazybones struct {
}

var hungryLazybones *HungryLazybones

var onc sync.Once

// GetHungryLazybonesInstance 获取单例实例
func GetHungryLazybonesInstance() *HungryLazybones {
	if hungryLazybones == nil {
		onc.Do(func() {
			hungryLazybones = new(HungryLazybones)
		})
	}
	return hungryLazybones
}
