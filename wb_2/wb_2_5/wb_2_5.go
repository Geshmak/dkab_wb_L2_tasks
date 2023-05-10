package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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

func foo(after int, before int, context int, count bool, ignoreCase bool, invert bool, fixed bool, lineNum bool, pattern string, files []string) {
	/*files = nil
	files = append(files, "kek.txt")
	lineNum = true
	pattern = "one"
	*/
	for _, fname := range files {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		var fileContents []string
		var resSlice []string
		counter := 0
		var strline []int

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {

			fileContents = append(fileContents, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		for i, cont := range fileContents {

			if (strings.Contains(strings.ToLower(cont), strings.ToLower(pattern)) && ignoreCase ||
				strings.Contains(cont, pattern) && !ignoreCase) != invert { //xor

				if fixed && ((strings.ToLower(cont) != strings.ToLower(pattern) && ignoreCase || cont != pattern && !ignoreCase) != invert) {
					continue
				}

				switch {
				case count:
					counter++

				case lineNum:
					strline = append(strline, i)

				case context != 0:
					start, finish := i-context, i+context+1
					inRange(&start, &finish, len(fileContents))
					resSlice = append(resSlice, fileContents[start:finish]...)

				case before != 0:
					start, finish := i-before, i+1
					inRange(&start, &finish, len(fileContents))
					resSlice = append(resSlice, fileContents[start:finish]...)

				case after != 0:
					start, finish := i, i+after+1
					inRange(&start, &finish, len(fileContents))
					resSlice = append(resSlice, fileContents[start:finish]...)

				default:
					resSlice = append(resSlice, cont)
				}
			}
		}
		strtmp := ""
		if len(files) > 1 {
			strtmp = fname + ": "
		}
		switch {
		case count:

			fmt.Println(strtmp + strconv.Itoa(counter))
		case lineNum:

			for _, num := range strline {
				fmt.Println(strtmp + strconv.Itoa(num))
			}
		default:
			for _, str := range resSlice {

				fmt.Println(strtmp + str)
			}
		}

	}
}
func main() {
	var after = flag.Int("A", 0, "печатать +N строк после совпадения")            //
	var before = flag.Int("B", 0, "печатать +N строк до совпадения")              //
	var context = flag.Int("C", 0, "печатать ±N строк вокруг совпадения")         //
	var count = flag.Bool("c", false, "количество строк")                         // +
	var ignoreCase = flag.Bool("i", false, "игнорировать регистр")                //+
	var invert = flag.Bool("v", false, "вместо совпадения, исключать")            //+
	var fixed = flag.Bool("F", false, "точное совпадение со строкой, не паттерн") //+
	var lineNum = flag.Bool("n", false, "печатать номер строки")                  // +

	flag.Parse()

	var pattern = flag.Args()[0]
	var files = flag.Args()[1:]

	/*var files []string
	files = append(files, "kek.txt")
	*lineNum = true
	pattern := "one"
	*/
	foo(*after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum, pattern, files)

}
