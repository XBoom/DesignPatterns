package _3_FactoryMethod

import (
	"reflect"
	"testing"
)

func TestBYDFactory_CreateCar(t *testing.T) {
	factory := new(BYDFactory)
	car := factory.CreateCar()
	if !reflect.DeepEqual(car.Name(), "BYD") {
		t.Fatalf("expect %s acture %s", "BYD", car.Name())
	}
}

func TestTeslaFactory_CreateCar(t *testing.T) {
	factory := new(TeslaFactory)
	car := factory.CreateCar()
	if !reflect.DeepEqual(car.Name(), "Tesla") {
		t.Fatalf("expect %s acture %s", "Tesla", car.Name())
	}
}
