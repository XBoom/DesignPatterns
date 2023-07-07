package _5_Builder

import (
	"reflect"
	"testing"
)

var ExpectBydCar = Car{"BYD Battery", "BYD Electric Motor", "BYD Electronic Control"}
var ExpectTeslaCar = Car{"Tesla Battery", "Tesla Electronic Motor", "Tesla Electronic Control"}

func TestBuilderBYD(t *testing.T) {
	//1. BYD图纸
	byd := new(BYD)
	//2. BYD工厂
	director := NewDirector(byd)
	//3. 生产BYD汽车
	director.Construct()
	//4. 客户拿到BYD汽车
	res := byd.GetCar()
	if reflect.DeepEqual(res, ExpectBydCar) {
		t.Fatalf("Builder1 fail expect 123 acture %s", res)
	}
}

func TestBuilderTesla(t *testing.T) {
	//1. Tesla图纸
	byd := new(Tesla)
	//2. Tesla工厂
	director := NewDirector(byd)
	//3. 生产Tesla汽车
	director.Construct()
	//4. 客户拿到Tesla汽车
	res := byd.GetCar()
	if reflect.DeepEqual(res, ExpectTeslaCar) {
		t.Fatalf("Builder1 fail expect 123 acture %s", res)
	}
}
