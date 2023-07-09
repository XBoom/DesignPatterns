package _7_Proxy

import "testing"

func TestStationProxy_sell(t *testing.T) {
	station := &Station{3}
	proxy := &StationProxy{station}
	station.Sell("A")
	proxy.Sell("B")
	proxy.Sell("C")
	proxy.Sell("D")
}
