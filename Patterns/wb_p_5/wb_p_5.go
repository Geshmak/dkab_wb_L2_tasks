package main

import "fmt"

type employee struct {
	name        string
	accessLevel int //0 - 3
}

type acesser interface {
	giveData(employee)
	setNextLvlGate(acesser)
}

type firstLevelGate struct {
	nextGate acesser
}

func (lvl1 *firstLevelGate) giveData(emp employee) {
	if emp.accessLevel >= 1 {
		if lvl1.nextGate != nil {
			lvl1.nextGate.giveData(emp)
		}
		fmt.Println("lvl1 secrets")
	} else {
		fmt.Println("---lvl1 access denied---")
	}
}
func (lvl1 *firstLevelGate) setNextLvlGate(nextGate acesser) {
	lvl1.nextGate = nextGate
}

type secondLevelGate struct {
	nextGate acesser
}

func (lvl2 *secondLevelGate) giveData(emp employee) {
	if emp.accessLevel >= 2 {
		if lvl2.nextGate != nil {
			lvl2.nextGate.giveData(emp)
		}
		fmt.Println("lvl2 secrets")
	} else {
		fmt.Println("---lvl2 access denied---")
	}
}
func (lvl2 *secondLevelGate) setNextLvlGate(nextGate acesser) {
	lvl2.nextGate = nextGate
}

type thirdLevelGate struct {
	nextGate acesser
}

func (lvl3 *thirdLevelGate) giveData(emp employee) {
	if emp.accessLevel >= 3 {
		if lvl3.nextGate != nil {
			lvl3.nextGate.giveData(emp)
		}
		fmt.Println("lvl3 secrets")
	} else {
		fmt.Println("---lvl3 access denied---")
	}
}
func (lvl3 *thirdLevelGate) setNextLvlGate(nextGate acesser) {
	lvl3.nextGate = nextGate
}

func create() *firstLevelGate {
	third := thirdLevelGate{}
	second := secondLevelGate{&third}
	first := firstLevelGate{&second}
	return &first
}

func main() {

	emp := employee{"Denis", 2}
	gate := create()
	gate.giveData(emp)
}
