package _3_FactoryMethod

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
	CreateCar() Car
}

type BYDFactory struct {
}

func (B BYDFactory) CreateCar() Car {
	return &BYD{}
}

var _ Factory = (*BYDFactory)(nil)

type TeslaFactory struct {
}

func (t TeslaFactory) CreateCar() Car {
	return &Tesla{}
}

var _ Factory = (*TeslaFactory)(nil)




