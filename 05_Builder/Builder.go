package _5_Builder

// Car 定义一个产品 汽车
type Car struct {
	battery           string //电池
	electricMotor     string //电机
	electronicControl string //电空
}

// Builder 抽象即按照这接口，定义其需要实现的方法
type Builder interface {
	BuildBattery()
	BuildElectricMotor()
	BuildElectronicControl()
	GetCar() *Car
}

// BYD 具体建造者
type BYD struct {
	car Car
}

func (b *BYD) BuildBattery() {
	b.car.battery = "BYD Battery"
}

func (b *BYD) BuildElectricMotor() {
	b.car.electricMotor = "BYD Electric Motor"
}

func (b *BYD) BuildElectronicControl() {
	b.car.electronicControl = "BYD Electronic Control"
}

func (b *BYD) GetCar() *Car {
	return &b.car
}

// 具体构建者实现抽象接口
var _ Builder = (*BYD)(nil)

// Tesla 具体建造者
type Tesla struct {
	car Car
}

func (t *Tesla) BuildBattery() {
	t.car.electronicControl = "Tesla Battery"
}

func (t *Tesla) BuildElectricMotor() {
	t.car.electronicControl = "Tesla Electronic Motor"
}

func (t *Tesla) BuildElectronicControl() {
	t.car.electronicControl = "Tesla Electronic Control"
}

func (t *Tesla) GetCar() *Car {
	return &t.car
}

var _ Builder = (*Tesla)(nil)

// Director 构建指挥者
type Director struct {
	builder Builder
}

// NewDirector 传入具体构建者，创建一个构建指挥者
func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.BuildBattery()
	d.builder.BuildElectricMotor()
	d.builder.BuildElectronicControl()
}
