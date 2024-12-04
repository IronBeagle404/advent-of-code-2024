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

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// For each line of the input, check wether the corresponding report is safe or unsafe and count the safe ones
	var safeReports int
	for scanner.Scan() {
		str := scanner.Text()
		if reportIsSafe(lineToArray(str)) {
			safeReports++
		}
	}

	file.Close()

	fmt.Printf("Safe reports : %v\n", safeReports)
}

// Convert a line of the input to an int array
func lineToArray(str string) []int {
	var intArray []int
	var tmpStr string

	for _, char := range str {
		if char != 32 {
			tmpStr += string(char)
		} else {
			i, _ := strconv.Atoi(tmpStr)
			intArray = append(intArray, i)
			tmpStr = ""
		}
	}
	if tmpStr != "" {
		i, _ := strconv.Atoi(tmpStr)
		intArray = append(intArray, i)
	}

	return intArray
}

// Mark report as safe or unsafe under given conditions
func reportIsSafe(report []int) bool {

	if !isArraySorted(report) {
		return false
	}

	var diff int
	for x, i := range report {
		if x != 0 {
			if i >= report[x-1] {
				diff = i - report[x-1]
			} else {
				diff = report[x-1] - i
			}

			if diff < 1 || diff > 3 {
				return false
			}
		}
	}

	return true
}

// Check if array is sorted in ascending or descending order
func isArraySorted(arr []int) bool {
	ascending := true
	descending := true

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			ascending = false
		}
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			descending = false
		}
	}

	return ascending || descending
}
