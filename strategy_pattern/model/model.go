package model

import "fmt"

type DuckBase struct {
	Color string `json:"color"`
	// 包含鸭子行为接口
	flayBehavior FlayBehavior
	quackBehavior QuackBehavior
}

// 所有鸭子都会的技能
func (d *DuckBase) Eat() {
	fmt.Println("duck is eat")
}

func (d *DuckBase) Display(color string) {
	fmt.Println(fmt.Sprintf("this is %s duck",color))
}

// 设置接口
func (d *DuckBase) SetFlayBehavior(f FlayBehavior) {
	d.flayBehavior = f
}

func (d *DuckBase) SetQuackBehavior(q QuackBehavior) {
	d.quackBehavior = q
}

// 实现
func (d *DuckBase) PerformFlay() {
	d.flayBehavior.Flay()
}

func (d *DuckBase) PerformQuack() {
	d.quackBehavior.Quack()
}

// 鸭子行为接口
type FlayBehavior interface {
	Flay()
}


type QuackBehavior interface {
	Quack()
}

