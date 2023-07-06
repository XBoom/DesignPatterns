package _1_Singleton

// 饿汉模式，直接在程序初始化的时候赋值

type HungrySingleton struct {
}

var hungrySingleton *HungrySingleton

func init() {
	hungrySingleton = new(HungrySingleton)
}

// GetHungryInstance 获取饿汉实例
func GetHungryInstance() *HungrySingleton {
	return hungrySingleton
}
