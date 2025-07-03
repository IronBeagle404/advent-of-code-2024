package main

import (
	"bufio"
	"fmt"
	"os"
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

	// Incomplete : Only works with 1-digit numbers
	var previousNbr, currentNbr, previousIncrDcr, currentIncrDcr, diff, safeReports int
	var validReport bool
	for scanner.Scan() {
		str := scanner.Text()
		validReport = true
		previousNbr, currentNbr, previousIncrDcr, currentIncrDcr = 0, 0, 0, 0
		for _, char := range str {
			if char != 32 {
				currentNbr = int(char - 48)
				if previousNbr != 0 {
					if previousNbr <= currentNbr {
						diff = currentNbr - previousNbr
						currentIncrDcr = 1
					} else if previousNbr > currentNbr {
						diff = previousNbr - currentNbr
						currentIncrDcr = -1
					}
					if (diff < 1 || diff > 3) || (previousIncrDcr != 0 && previousIncrDcr != currentIncrDcr) {
						validReport = false
						break
					}
				}
				previousNbr = currentNbr
				previousIncrDcr = currentIncrDcr
			}
		}
		if validReport {
			safeReports++
		}
	}
	file.Close()

	fmt.Printf("Safe reports : %v\n", safeReports)
}
