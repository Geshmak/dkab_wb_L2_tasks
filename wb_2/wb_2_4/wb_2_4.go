package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagram(strsl []string) map[string][]string {
	tmpMap := make(map[string]string)
	resMap := make(map[string][]string)

	//sort.Slice(strsl, func(i, j int) bool { return strsl[i] < strsl[j] })

	for _, str := range strsl {
		str = strings.ToLower(str)
		strk := getKey(str)

		if val, boo := tmpMap[strk]; !boo {
			tmpMap[strk] = str
			resMap[str] = append(resMap[str], str)
		} else {
			resMap[val] = append(resMap[val], str)
		}
	}
	for key, val := range resMap {
		sort.Slice(val, func(i, j int) bool { return val[i] < val[j] })
		if len(resMap[key]) == 1 {
			delete(resMap, key)
		}
	}
	return resMap
}

func getKey(str1 string) string {
	chslice1 := strings.Split(strings.ToLower(str1), "")
	sort.Slice(chslice1, func(i, j int) bool { return chslice1[i] < chslice1[j] })

	return strings.Join(chslice1, "")
}

func main() {
	arr := []string{"пятак", "пятак", "пятка", "Яптка", "апятк", "тяпка", "листок", "ислток", "слиток", "столик", "полс"}
	res := anagram(arr)
	fmt.Println(res)
}
