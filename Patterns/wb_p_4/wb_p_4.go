package main

import "fmt"

type base struct {
	soildersNum int
}

func NewBase(num int) *base {
	return &base{num}
}
func (b *base) paintGrass(num int) command {
	return &paintGrassCommand{num, b}
}
func (b *base) paintGravel(num int) command {
	return &paintGravelCommand{num, b}
}
func (b *base) paintSnow(num int) command {
	return &paintSnowCommand{num, b}
}

type command interface {
	execute()
}

type paintGrassCommand struct {
	num  int
	base *base
}

func (c paintGrassCommand) execute() {
	c.base.paintGrass(c.num)
	fmt.Println("газон покрашен", c.num)
}

type paintGravelCommand struct {
	num  int
	base *base
}

func (c paintGravelCommand) execute() {
	c.base.paintGravel(c.num)
	fmt.Println("щебень покрашен", c.num)
}

type paintSnowCommand struct {
	num  int
	base *base
}

func (c paintSnowCommand) execute() {
	c.base.paintSnow(c.num)
	fmt.Println("снег покрашен", c.num)
}

type officer struct {
	commands []command
}

func (e *officer) executeCommands() {
	for _, c := range e.commands {
		c.execute()
	}
}

func main() {
	b := NewBase(50)

	commands := []command{
		b.paintGrass(4),
		b.paintGravel(2),
		b.paintSnow(6),
		b.paintGravel(2),
		b.paintGrass(20),
	}

	for _, c := range commands {
		c.execute()
	}

	officer := officer{commands}
	officer.executeCommands()
}
