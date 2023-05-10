package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isdigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func giveInterval(runes []rune) func() (int, int) {
	start, finish := 0, 0
	return func() (int, int) {
		startChanged := false
		start = finish + 1
		for i := start; i < len(runes); i++ {
			if isdigit(runes[i]) && !startChanged {
				start = i
				startChanged = true
			}

			if !isdigit(runes[i]) && startChanged {
				finish = i - 1
				return start, finish
			}
		}
		if startChanged {
			finish = len(runes) - 1
			return start, finish
		}
		return -1, -1

	}
}

func unpack(str string) string {
	if str == "" {
		return ""
	}

	runes := []rune(str)
	if isdigit(runes[0]) {
		return ""
	}

	res := ""
	giver := giveInterval(runes)
	start, finish := giver()
	j := 0

	for start != -1 {
		res += string(runes[j : start-1])
		num, _ := strconv.Atoi(string(runes[start : finish+1]))
		res += strings.Repeat(string(runes[start-1]), num)

		j = finish + 1
		start, finish = giver()
	}

	if j < len(runes) {
		res += string(runes[j:])
	}

	return res
	/*
		tmp := ""

		for i := 0; i < len(runes); i++ {

			fmt.Println(string(runes[i]))
			if !isdigit(runes[i]) {
				res += string(runes[i])
			} else if res != "" {
				j := i - 1

				for isdigit(runes[i]) {
					tmp += string(runes[i])
					i++
				}

				num, _ := strconv.Atoi(tmp)
				res += strings.Repeat(string(runes[j]), num-1)
				tmp = ""
				i--
			} else {
				return "некорректная строка"
			}
		}
		return res*/
}

func main() {
	str := "a4bc2d11e3"

	fmt.Println(unpack(str))

	fmt.Println(unpack("abcd"))
	fmt.Println(unpack("45"))
	fmt.Println(unpack("d5e"))
	fmt.Println(unpack(""))
}
