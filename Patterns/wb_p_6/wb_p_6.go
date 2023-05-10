package main

import (
	"fmt"
	"strings"
)

type tranporter interface {
	getName() string
}

type transport struct {
	name string
}

func (t transport) getName() string {
	return t.name
}

type ship struct {
	transport
	port string
}

type plane struct {
	transport
	airport string
}

type train struct {
	transport
	trainstation string
}

func newShip() tranporter {
	return &ship{transport: transport{name: "ship"}, port: "Sbp"}
}

func newPlane() tranporter {
	return &plane{transport: transport{name: "plane"}, airport: "Domodedovo"}
}

func newTrain() tranporter {
	return &train{transport: transport{name: "train"}, trainstation: "Paveletskiy"}
}

func getTransport(str string) (bool, tranporter) {
	switch strings.ToLower(str) {
	case "ship":
		return true, newShip()
	case "train":
		return true, newTrain()
	case "plane":
		return true, newPlane()
	default:
		fmt.Println("wrong format")
		return false, nil
	}

}
func main() {

	if boo, val := getTransport("ship"); boo {
		fmt.Println(val.getName())
	}
	if boo, val := getTransport("plane"); boo {
		fmt.Println(val.getName())
	}
	if boo, val := getTransport("sheeesh"); boo {
		fmt.Println(val.getName())
	}

}
