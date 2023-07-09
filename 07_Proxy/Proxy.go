package _7_Proxy

import "fmt"

// Seller 售票
type Seller interface {
	Sell(name string)
}

// Station 火车站
type Station struct {
	stock int //库存
}

func (station *Station) Sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("%s purchased 1 ticket, %d remaining\n", name, station.stock)
	} else {
		fmt.Println("tickets are sold out")
	}
}

var _ Seller = (*Station)(nil)

// StationProxy 火车代理点
type StationProxy struct {
	station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) Sell(name string) {
	if proxy.station.stock > 0 {
		proxy.station.stock--
		fmt.Printf("%s purchased 1 ticket, %d remaining\n", name, proxy.station.stock)
	} else {
		fmt.Println("tickets are sold out")
	}
}

var _ Seller = (*StationProxy)(nil)
