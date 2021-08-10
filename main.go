package main

import (
	"fmt"
	"strings"
)

type KeyValue struct {
	Key   string
	Value int
}

func main() {
	str := "ccccabbaabcdabeeee"
	strMap := make(map[string]int)

	for i := 0; i < len(str); i++ {
		_, ok := strMap[string(str[i])]
		if ok {
			strMap[string(str[i])] += 1
		} else {
			strMap[string(str[i])] = 1

		}
	}

	var max int
	var tempChar string
	finalMaxMap := make(map[string]int)
	for k, v := range strMap {
		if max == v {
			finalMaxMap[k] = v
			index1 := strings.Index(str, k)
			index2 := strings.Index(str, tempChar)
			fmt.Println("index", index1, index2)
			if index1 < index2 {
				delete(finalMaxMap, tempChar)
			} else {
				delete(finalMaxMap, k)
			}
		} else if strMap[k] > max {
			if tempChar != "" {
				delete(finalMaxMap, tempChar)
			}
			finalMaxMap[k] = v
			tempChar = k
			max = v

		}

	}

	fmt.Println("final", finalMaxMap)
}
