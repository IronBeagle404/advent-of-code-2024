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

	// For each line of the input, check if the report is safe or can it be fixed with the problem dampener
	var safeReports int
	for scanner.Scan() {
		str := scanner.Text()
		if reportIsSafe(lineToArray(str)) || problemDampener(lineToArray(str)) {
			safeReports++
		}
	}

	file.Close()

	fmt.Printf("Safe reports : %v\n", safeReports)
}

// basic dampener that check if any number is removable from the report to make it correct
func problemDampener(report []int) bool {
	canBeResolved := false
	var tmpReport []int
	for x := range report {
		if x == 0 {
			if reportIsSafe(report[x+1:]) {
				canBeResolved = true
				break
			}
		} else if x == len(report)-1 {
			if reportIsSafe(report[:x]) {
				canBeResolved = true
				break
			}
		} else {
			tmpReport = append(tmpReport, report[:x]...)
			tmpReport = append(tmpReport, report[x+1:]...)
			if reportIsSafe(tmpReport) {
				canBeResolved = true
				break
			}
			tmpReport = nil
		}
	}
	return canBeResolved
}

// Broken dampener attempt, made with the idea that it should try to remove only incorrect numbers
// Overcomplicated and buggy as hell, would have liked to make it work though
// (yes i left my crappy testing comments & all the bullshit in it, I just like the idea of it still being here just to show that I tried to do something smart lmao)
//
// Original function comment :
//// Works the same way than the reportIsSafe func, but every time a number is incorect it checks if the report would be correct without said number
//
// func problemDampener(report []int) bool {
// 	canBeResolved := false
// 	var testReport []int
// 	var diff int
// 	//ordered := true
// 	//fmt.Printf("Report before check is %v\n", report)
// 	for x, i := range report {
// 		ordered := true
// 		if x != 0 {
// 			if x >= 1 && x <= len(report)-2 {
// 				if !isArraySorted(report[x-1 : x+2]) {
// 					// fmt.Printf("unsorted extract is %v\n", report[x-1:x+2])
// 					// fmt.Printf("faulty data is %v\n", report[x])
// 					ordered = false
// 				}
// 			}
// 			if i >= report[x-1] {
// 				diff = i - report[x-1]
// 			} else {
// 				diff = report[x-1] - i
// 			}
// 			// if !ordered {
// 			// 	fmt.Printf("faulty data is %v\n", report[x])
// 			// }
// 			if (diff < 1 || diff > 3) || !ordered {
// 				//fmt.Printf("faulty data is %v\n", report[x])
// 				if x == len(report)-1 {
// 					testReport = report[0 : x-1]
// 					//fmt.Printf("testreport is %v\n", testReport)
// 				} else {
// 					testReport = report[0:x]
// 					testReport = append(testReport, report[x+1:]...)
// 					//fmt.Printf("test report is %v\n", testReport)
// 				}
// 				// testReport = report[0:x]
// 				// testReport = append(testReport, report[x+1])
// 				//fmt.Printf("faulty data is %v\n", report[x])
// 				if reportIsSafe(testReport) {
// 					//fmt.Printf("%v was not safe but %v is\n\n", report, testReport)
// 					canBeResolved = true
// 					break
// 				} else {
// 					//fmt.Printf("Report after check is %v\n", report)
// 					//fmt.Printf("%v cant be corrected\n", report[x])
// 				}
// 			}
// 		}
// 	}
// 	return canBeResolved
// }

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
// // changed the sorting check to something a bit more precise
func reportIsSafe(report []int) bool {

	ordered := true

	var diff int
	for x, i := range report {

		if x >= 1 && x <= len(report)-2 {
			if !isArraySorted(report[x-1 : x+2]) {
				ordered = false
			}
		}

		if x != 0 {
			if i >= report[x-1] {
				diff = i - report[x-1]
			} else {
				diff = report[x-1] - i
			}

			if (diff < 1 || diff > 3) || !ordered {
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
