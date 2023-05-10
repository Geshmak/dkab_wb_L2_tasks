package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func sortf(fname string, reverse bool, number bool, nodup bool, field string) {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var fileContents []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContents = append(fileContents, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if number {
		var numContents []int
		for _, sToNum := range fileContents {
			num, err := strconv.Atoi(sToNum)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			numContents = append(numContents, num)
		}

		if reverse {
			//reverse
			sort.Slice(numContents, func(i, j int) bool {
				return numContents[i] > numContents[j]
			})
		} else {
			//default
			sort.Slice(numContents, func(i, j int) bool {
				return numContents[i] < numContents[j]
			})
		}

		if nodup {
			//no duplicates
			numContents = removeDuplicate(numContents)
		}
		fileContents = nil
		for _, nToStr := range numContents {
			str := strconv.Itoa(nToStr)
			fileContents = append(fileContents, str)
		}
	} else {

		if field != "" {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if reverse {
				//reverse
				sort.Slice(fileContents, func(i, j int) bool {
					return strings.Split(fileContents[i], " ")[num] > strings.Split(fileContents[j], " ")[num]
				})
			} else {
				//default
				sort.Slice(fileContents, func(i, j int) bool {
					return strings.Split(fileContents[i], " ")[num] < strings.Split(fileContents[j], " ")[num]
				})
			}
		} else {
			if reverse {
				//reverse
				sort.Slice(fileContents, func(i, j int) bool {
					return fileContents[i] > fileContents[j]
				})
			} else {
				//default
				sort.Slice(fileContents, func(i, j int) bool {
					return fileContents[i] < fileContents[j]
				})
			}
		}

		if nodup {
			//no duplicates
			fileContents = removeDuplicate(fileContents)
		}
	}
	newfile, err := os.Create("1" + fname)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer newfile.Close()
	for _, str := range fileContents {
		newfile.WriteString(str + "\n")
	}

}

func main() {

	fields := flag.String("k", "", "указание колонки для сортировки")
	reverse := flag.Bool("r", false, "сортировка в обратном порядке")
	number := flag.Bool("n", false, "сортировать по числовому значению")
	nodup := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	if *reverse {
		fmt.Println("1")
	}
	if *number {
		fmt.Println("2")
	}
	if *nodup {
		fmt.Println("3")
	}
	//args := flag.Args()

	sortf("kek.txt", *reverse, *number, *nodup, *fields)

}
