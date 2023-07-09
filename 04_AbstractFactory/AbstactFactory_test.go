package _4_AbstractFactory

import (
	"reflect"
	"testing"
)

var expectBYD = "BYD"
var expectTesla = "Tesla"

func TestShangHaiFactory(t *testing.T) {
	factory := new(ShangHaiFactory)
	byd := factory.CreateBYD()
	if !reflect.DeepEqual(byd.Name(), expectBYD) {
		t.Fatalf("expect %s acture %s", expectBYD, byd.Name())
	}
	tesla := factory.CreateTesla()
	if !reflect.DeepEqual(tesla.Name(), expectTesla) {
		t.Fatalf("expect %s acture %s", expectTesla, tesla.Name())
	}

}

func TestWuhanFactory(t *testing.T) {
	factory := new(WuhanFactory)
	byd := factory.CreateBYD()
	if !reflect.DeepEqual(byd.Name(), expectBYD) {
		t.Fatalf("expect %s acture %s", expectBYD, byd.Name())
	}
	tesla := factory.CreateTesla()
	if !reflect.DeepEqual(tesla.Name(), expectTesla) {
		t.Fatalf("expect %s acture %s", expectTesla, tesla.Name())
	}
}
