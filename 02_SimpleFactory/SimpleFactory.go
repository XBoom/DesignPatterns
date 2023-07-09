package _2_SimpleFactory

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

type Factory struct {
}

func (f *Factory) CreateCar(typ string) Car {
	switch typ {
	case "byd":
		return &BYD{}
	case "tesla":
		return &Tesla{}
	default:
		panic("unknown type")
	}
}
