package _0_Decorator

import (
	"strings"
	"testing"
)

var ExpectBattery = "Condensed Battery"
var ExpectBYD = "BYD use Condensed Battery run"
var ExpectTesla = "Tesla use Condensed Battery run"

func TestBYD_Run(t *testing.T) {
	battery := new(BatteryComponent)
	res := battery.Name()
	if strings.Compare(res, ExpectBattery) != 0 {
		t.Fatalf("Eoncrete fail expect %s acture %s", ExpectBattery, res)
	}

	byd := &BYD{battery}
	res = byd.Run()
	if strings.Compare(res, ExpectBYD) != 0 {
		t.Fatalf("Eoncrete fail expect %s acture %s", ExpectBYD, res)
	}
}

func TestTesla_Run(t *testing.T) {
	battery := new(BatteryComponent)
	res := battery.Name()
	if strings.Compare(res, ExpectBattery) != 0 {
		t.Fatalf("Eoncrete fail expect %s acture %s", ExpectBattery, res)
	}

	tesla := &Tesla{battery}
	res = tesla.Run()
	if strings.Compare(res, ExpectTesla) != 0 {
		t.Fatalf("Eoncrete fail expect %s acture %s", ExpectTesla, res)
	}
}
