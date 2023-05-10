package main

import "fmt"

type State interface {
	pressButton()
}

type Computer struct {
	currState State
}

func (c *Computer) pressButton() {
	c.currState.pressButton()
}

func (c *Computer) changeState(st State) {
	c.currState = st
}

func newComputer() *Computer {
	computer := Computer{}
	computer.changeState(&TurnedOff{computer: &computer})
	return &computer
}

// TurnedOff State
type TurnedOff struct {
	computer *Computer
}

func (t *TurnedOff) pressButton() {
	fmt.Println("Nothing")
	t.computer.changeState(&TurnedOn{computer: t.computer})
}

// Asleep state
type Asleep struct {
	computer *Computer
}

func (a *Asleep) pressButton() {
	fmt.Println("Wakes up the computer")
	a.computer.changeState(&TurnedOn{computer: a.computer})

}

// TurnedOn State
type TurnedOn struct {
	computer *Computer
}

func (t *TurnedOn) pressButton() {
	fmt.Println("do something")
	t.computer.changeState(&Asleep{computer: t.computer})
}

func main() {
	computer := newComputer()
	for i := 0; i < 10; i++ {
		computer.pressButton()
	}
}
