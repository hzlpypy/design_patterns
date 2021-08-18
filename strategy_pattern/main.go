package main

// 策略模式
import (
	"fmt"
	"github.com/hzlpypy/design_patterns/strategy_pattern/model"
)

// 绿鸭子：会飞，会叫
type GreenDuck struct {
	*model.DuckBase
}

type GreenDuckBehavior struct {
	*GreenDuck
}

func (g *GreenDuckBehavior)Flay()  {
	fmt.Println("GreenDuck:Flay")
}

func (g *GreenDuckBehavior)Quack()  {
	fmt.Println("GreenDuck:Quack")
}

// 橡皮鸭子：不会飞，不会叫
type RubberDuck struct {
	*model.DuckBase
}

func main() {
	gd := &GreenDuck{DuckBase:&model.DuckBase{
		Color: "green",
	}}
	gdb := &GreenDuckBehavior{gd}
	gd.SetFlayBehavior(gdb)
	gd.SetQuackBehavior(gdb)
	rd := &RubberDuck{DuckBase:&model.DuckBase{
		Color: "gray",
	}}
	gd.Display(gd.Color)
	rd.Display(rd.Color)
	gd.PerformFlay()
	gd.PerformQuack()

}