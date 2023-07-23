package _5_Template

import "fmt"

// Template 模板流程
type Template interface {
	Step1()
	Step2()
	Step3()
}

// 实现模板方法类似抽象类

type ConcreteTemplate struct{}

func (t *ConcreteTemplate) Step1() {
	fmt.Println("ConcreteTemplate:step1")
}

func (t *ConcreteTemplate) Step2() {
	fmt.Println("ConcreteTemplate:step2")
}

func (t *ConcreteTemplate) Step3() {
	fmt.Println("ConcreteTemplate:step3")
}

func GetResult(t Template) {
	t.Step1()
	t.Step2()
	t.Step3()
}

type ConcreteTemplateA struct {
	ConcreteTemplate
}

func (t *ConcreteTemplateA) Step1() {
	fmt.Println("ConcreteTemplateA:step1")
}

type ConcreteTemplateB struct {
	ConcreteTemplate
}

func (t *ConcreteTemplateB) Step3() {
	fmt.Println("ConcreteTemplateB:step3")
}
