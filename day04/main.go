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

// Search every occurence of XMAS/SAMX in the input
func FindXmas(matrix [][]rune) int {
	var xmasFound int
	for x, line := range matrix {
		for y, char := range line {
			// Horizontal
			if (y <= len(line)-4) && (string(char)+string(matrix[x][y+1])+string(matrix[x][y+2])+string(matrix[x][y+3]) == "XMAS" || string(char)+string(matrix[x][y+1])+string(matrix[x][y+2])+string(matrix[x][y+3]) == "SAMX") {
				xmasFound++
			}
			// Vertical
			if (x <= len(matrix)-4) && (string(char)+string(matrix[x+1][y])+string(matrix[x+2][y])+string(matrix[x+3][y]) == "XMAS" || string(char)+string(matrix[x+1][y])+string(matrix[x+2][y])+string(matrix[x+3][y]) == "SAMX") {
				xmasFound++
			}
			// Diagonal right
			if (y <= len(line)-4 && x <= len(matrix)-4) && (string(char)+string(matrix[x+1][y+1])+string(matrix[x+2][y+2])+string(matrix[x+3][y+3]) == "XMAS" || string(char)+string(matrix[x+1][y+1])+string(matrix[x+2][y+2])+string(matrix[x+3][y+3]) == "SAMX") {
				xmasFound++
			}
			// Diagonal left
			if (y >= 3 && x <= len(matrix)-4) && (string(char)+string(matrix[x+1][y-1])+string(matrix[x+2][y-2])+string(matrix[x+3][y-3]) == "XMAS" || (string(char)+string(matrix[x+1][y-1])+string(matrix[x+2][y-2])+string(matrix[x+3][y-3]) == "SAMX")) {
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
