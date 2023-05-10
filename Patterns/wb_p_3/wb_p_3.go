package main

import (
	"fmt"
	"strconv"
)

// throughput
type square struct {
	throughput int
	size       int
	name       string
}

func (s square) showInfo() string {
	return s.name + ": size - " + strconv.Itoa(s.size) + " throughput - " + strconv.Itoa(s.throughput)
}
func (s square) accept(v visitor) {
	v.visitSquare(s)
}

type museum struct {
	tickets         int
	artInstalations int
	name            string
	location        string
	ticketCost      int
}

func (m museum) showInfo() string {
	return m.name + " " + m.location + ": artInstalations - " + strconv.Itoa(m.artInstalations)
}
func (m museum) accept(v visitor) {
	v.visitMuseum(m)
}

type visit struct {
	money int
}
type visitor interface {
	visitSquare(square)
	visitMuseum(museum)
}

type countPeople struct {
	days int
}

func (c countPeople) visitSquare(s square) {
	fmt.Println(s.throughput * c.days)
}
func (c countPeople) visitMuseum(m museum) {
	fmt.Println(m.tickets * c.days)
}

func main() {
	mus := museum{125, 10, "art museum", "Moscow", 50}
	squar := square{2000, 300, "main square"}

	fmt.Println(mus.showInfo())
	fmt.Println(squar.showInfo())

	mus.accept(countPeople{5})
	squar.accept(countPeople{10})
}
