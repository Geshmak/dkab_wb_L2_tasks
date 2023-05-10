package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func inRange(start *int, finish *int, len int) {
	if *start < 0 {
		*start = 0
	}
	if *finish >= len {
		*finish = len
	}
}
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func foo(fields string, delimiter string, separated bool) {
	reader := bufio.NewReader(os.Stdin)
	var (
		start      int
		finish     int
		fl         bool
		numSlice   []int
		inputSlice [][]string
	)

	for {
		text, err := reader.ReadString('\n')
		if err == io.EOF || strings.Contains(text, "---stop---") {
			fmt.Println("\n-------------------------------")
			break
		}
		if !strings.Contains(text, delimiter) && separated {
			continue
		}
		newStr := strings.Split(text, delimiter)
		inputSlice = append(inputSlice, newStr)
	}

	switch {
	case strings.IndexRune(fields, '-') != -1:
		i := strings.IndexRune(fields, '-')
		fslice := strings.Split(fields, "-")
		start = 0
		finish = 0

		if fslice[1] == "" && i == 0 {
			start, _ = strconv.Atoi(fslice[0])
			finish = 0
		} else if fslice[1] == "" && i == 0 {
			start = 0
			finish, _ = strconv.Atoi(fslice[0])
		} else {
			start, _ = strconv.Atoi(fslice[0])
			finish, _ = strconv.Atoi(fslice[1])
		}

	default:
		tmpstr := strings.Split(fields, ",")

		for _, num := range tmpstr {
			val, _ := strconv.Atoi(num)
			numSlice = append(numSlice, val)
		}
		start, finish = -1, 0
	}

	if finish == 0 {
		fl = true
	}
	for i := 0; i < len(inputSlice); i++ {
		strSlice := inputSlice[i]
		if fl {
			finish = len(strSlice) - 1
		}

		if start == -1 {
			for i, word := range strSlice {
				if contains(numSlice, i) {
					fmt.Print(word + " ")
				}

			}
		} else {

			inRange(&start, &finish, len(strSlice))
			for _, word := range strSlice[start:finish] {
				fmt.Print(word + " ")
			}
		}
		fmt.Print("\n")
	}

}

func main() {
	var fields = flag.String("f", "", "выбрать поля (колонки)")
	var delimiter = flag.String("d", "\t", "использовать другой разделитель")
	var separated = flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	foo(*fields, *delimiter, *separated)

}
