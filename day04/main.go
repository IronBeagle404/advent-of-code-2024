package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputStr := string(input)
	matrix := StrToRuneMatrix(inputStr)

	fmt.Printf("XMAS found : %v\n", FindXmas(matrix))
}

// Search every occurence of a MAS/SAM cross in the input
func FindXmas(matrix [][]rune) int {
	var xmasFound int
	for x, line := range matrix {
		for y, char := range line {
			if (x >= 1 && x <= len(matrix)-2) && (y >= 1 && y <= len(line)-2) && (string(matrix[x-1][y-1])+string(char)+string(matrix[x+1][y+1]) == "MAS") && (string(matrix[x+1][y-1])+string(char)+string(matrix[x-1][y+1]) == "MAS") {
				xmasFound++
			}
			if (x >= 1 && x <= len(matrix)-2) && (y >= 1 && y <= len(line)-2) && (string(matrix[x-1][y-1])+string(char)+string(matrix[x+1][y+1]) == "SAM") && (string(matrix[x+1][y-1])+string(char)+string(matrix[x-1][y+1]) == "SAM") {
				xmasFound++
			}
			if (x >= 1 && x <= len(matrix)-2) && (y >= 1 && y <= len(line)-2) && (string(matrix[x-1][y-1])+string(char)+string(matrix[x+1][y+1]) == "SAM") && (string(matrix[x-1][y+1])+string(char)+string(matrix[x+1][y-1]) == "SAM") {
				xmasFound++
			}
			if (x >= 1 && x <= len(matrix)-2) && (y >= 1 && y <= len(line)-2) && (string(matrix[x-1][y-1])+string(char)+string(matrix[x+1][y+1]) == "MAS") && (string(matrix[x-1][y+1])+string(char)+string(matrix[x+1][y-1]) == "MAS") {
				xmasFound++
			}
		}
	}
	return xmasFound
}

func StrToRuneMatrix(str string) [][]rune {
	var matrix [][]rune
	var runeArr []rune

	for _, char := range str {
		if char != 10 {
			runeArr = append(runeArr, rune(char))
		} else {
			matrix = append(matrix, runeArr)
			runeArr = nil
		}
	}

	if runeArr != nil {
		matrix = append(matrix, runeArr)
		runeArr = nil
	}

	return matrix
}
