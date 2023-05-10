package main

import (
	"fmt"
	"strconv"
)

type animalBuild struct {
	legs int
	tail bool
	roar string
}

type animalBuilder interface {
	makeLegs(int) animalBuilder
	makeTail(bool) animalBuilder
	makeRoar(string) animalBuilder
	Build() animaler
}

func New() animalBuilder {
	return &animalBuild{}
}
func (a *animalBuild) makeLegs(num int) animalBuilder {
	a.legs = num
	return a
}
func (a *animalBuild) makeTail(flag bool) animalBuilder {
	a.tail = flag
	return a
}
func (a *animalBuild) makeRoar(str string) animalBuilder {
	a.roar = str
	return a
}
func (a *animalBuild) Build() animaler {
	return &animal{
		legs: a.legs,
		tail: a.tail,
		roar: a.roar,
	}
}

type animaler interface {
	produceRoar() string
	showYourself() string
}

type animal struct {
	legs int
	tail bool
	roar string
}

func (a *animal) produceRoar() string {
	return a.roar
}
func (a *animal) showYourself() string {
	res := ""
	if a.tail {
		res += "big tail and "
	}
	res += strconv.Itoa(a.legs) + " legs"
	return res
}

func main() {
	builder := New()
	cat := builder.makeLegs(4).makeRoar("mew").makeTail(true).Build()
	fmt.Println(cat.produceRoar())
	fmt.Println(cat.showYourself())

	human := builder.makeLegs(2).makeRoar("sheeesh").makeTail(false).Build()
	fmt.Println(human.produceRoar())
	fmt.Println(human.showYourself())
}
