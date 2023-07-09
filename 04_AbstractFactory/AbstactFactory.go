package _4_AbstractFactory

type Car interface {
	Name() string
}

type BYD struct {
}

func (B BYD) Name() string {
	return "BYD"
}

var _ Car = (*BYD)(nil)

type Tesla struct {
}

func (t Tesla) Name() string {
	return "Tesla"
}

var _ Car = (*Tesla)(nil)

type Factory interface {
	CreateBYD() Car
	CreateTesla() Car
}

type WuhanFactory struct {
}

func (w WuhanFactory) CreateBYD() Car {
	return &BYD{}
}

func (w WuhanFactory) CreateTesla() Car {
	return &Tesla{}
}

var _ Factory = (*WuhanFactory)(nil)

type ShangHaiFactory struct {
}

func (s ShangHaiFactory) CreateBYD() Car {
	return &BYD{}
}

func (s ShangHaiFactory) CreateTesla() Car {
	return &Tesla{}
}

var _ Factory = (*ShangHaiFactory)(nil)