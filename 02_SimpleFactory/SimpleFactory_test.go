package _2_SimpleFactory

import (
	"reflect"
	"testing"
)

var expectByd = "BYD"
var expectTesla = "Tesla"

func TestFactory_CreateCar(t *testing.T) {
	factory := new(Factory)
	car1 := factory.CreateCar("byd")
	if !reflect.DeepEqual(car1.Name(), expectByd) {
		t.Fatalf("expect %s acture %s", expectByd, car1.Name())
	}

	car2 := factory.CreateCar("tesla")
	if !reflect.DeepEqual(car2.Name(), expectTesla) {
		t.Fatalf("expect %s acture %s", expectTesla, car2.Name())
	}
}
