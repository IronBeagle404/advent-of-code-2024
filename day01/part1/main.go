package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read input & put the values of each line into two arrays
	var list1, list2 []int
	var tmpStr1, tmpStr2 string
	var data1, data2 int
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
	for scanner.Scan() {
		str := scanner.Text()
		afterSpace := false
		tmpStr1, tmpStr2 = "", ""
		for _, char := range str {
			if char == 32 {
				afterSpace = true
			}
			if !afterSpace {
				tmpStr1 += string(char)
			} else if char != 32 {
				tmpStr2 += string(char)
			}
		}
		data1, _ = strconv.Atoi(tmpStr1)
		data2, _ = strconv.Atoi(tmpStr2)
		list1 = append(list1, data1)
		list2 = append(list2, data2)
	}

	file.Close()

	// Sort lists
	bubbleSort(list1)
	bubbleSort(list2)

	// Calculate difference between list items
	res := 0
	var diff int
	for x := 0; x < len(list1); x++ {
		if list1[x] > list2[x] {
			diff = list1[x] - list2[x]
		} else {
			diff = list2[x] - list1[x]
		}
		res += diff
	}

	// Print result
	fmt.Println(res)
}

func bubbleSort(array []int) {
	arrayLen := len(array) - 1
	var swapped bool
	for x := 0; x < arrayLen; x++ {
		swapped = false
		for y := 0; y < arrayLen-x; y++ {
			if array[y] > array[y+1] {
				tmp := array[y]
				array[y] = array[y+1]
				array[y+1] = tmp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
