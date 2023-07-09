package _3_FactoryMethod

import (
	"reflect"
	"testing"
)

var expectBYD = "BYD"
var expectTesla = "Tesla"

func TestBYDFactory_CreateCar(t *testing.T) {
	factory := new(BYDFactory)
	car := factory.CreateCar()
	if !reflect.DeepEqual(car.Name(), expectBYD) {
		t.Fatalf("expect %s acture %s", expectBYD, car.Name())
	}
}

func TestTeslaFactory_CreateCar(t *testing.T) {
	factory := new(TeslaFactory)
	car := factory.CreateCar()
	if !reflect.DeepEqual(car.Name(), expectTesla) {
		t.Fatalf("expect %s acture %s", expectTesla, car.Name())
	}
}
