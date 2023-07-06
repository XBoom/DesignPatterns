package _1_Singleton

import "sync"

type LazybonesSingleton struct {
}

var once sync.Once
var lazybonesSingleton *LazybonesSingleton

// GetLazybonesInstance 获取懒汉实例
func GetLazybonesInstance() *LazybonesSingleton {
	once.Do(func() {
		lazybonesSingleton = new(LazybonesSingleton)
	})
	return lazybonesSingleton
}
