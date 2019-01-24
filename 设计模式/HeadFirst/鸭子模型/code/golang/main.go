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
	fb FlyBehavior
}

func (d *Duck) performFly() {
	d.fb.fly()
}

func main() {
	fmt.Println("duck")
	duck := &Duck{
		fb: &FlyNoWay{},
	}
	duck.performFly()
	duck.fb = &FlyWithWings{}
	duck.performFly()
}
