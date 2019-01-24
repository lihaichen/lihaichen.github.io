package main

import "fmt"

type FlyBehavior interface {
	fly()
}

type FlyWithWings struct {
}

func (m *FlyWithWings) fly() {
	fmt.Println("flyWithWings")
}

type FlyNoWay struct {
}

func (m *FlyNoWay) fly() {
	fmt.Println("FlyNoWay")
}

type Duck struct {
	f FlyBehavior
}

func (m *Duck) performFly() {
	m.f.fly()
}

func main() {
	fmt.Println("duck")
	duck := &Duck{
		f: &FlyNoWay{},
	}
	duck.performFly()
	duck.f = &FlyWithWings{}
	duck.performFly()
}
