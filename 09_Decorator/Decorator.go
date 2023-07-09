package _9_Decorator

// Battery 抽象构件，给这个对象动态地添加职责
type Battery interface {
	Name() string
}

// BatteryComponent 具体构件
type BatteryComponent struct{}

func (c *BatteryComponent) Name() string {
	return "Condensed Battery"
}

// Car 装饰者
type Car interface {
	Run() string
}

// BYD 具体装饰者 BYD
type BYD struct {
	battery Battery
}

func (d *BYD) Run() string {
	return "BYD use " + d.battery.Name() + " run"
}

// Tesla 具体装饰者 Tesla
type Tesla struct {
	battery Battery
}

func (t *Tesla) Run() string {
	return "Tesla use " + t.battery.Name() + " run"
}
