package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	contentStr := string(content)

	// Define a regex corresponding to the correct instruction syntax
	re := regexp.MustCompile(`mul\(\d*,\d*\)`)

	// Find all instances of the regex in the input
	found := re.FindAllString(contentStr, -1)
	if found == nil {
		fmt.Println("No match found")
		return
	}

	// Define a regex for the numbers
	nbrRe, nbrErr := regexp.Compile(`\d+`)
	if nbrErr != nil {
		log.Fatal(nbrErr)
	}

	// Identify the two numbers in each instance, multiply them together and add it to the total sum
	var nbr1, nbr2, instanceRes, totalRes int
	for _, instance := range found {
		nbrFound := nbrRe.FindAllString(instance, -1)
		if nbrFound != nil {
			nbr1, _ = strconv.Atoi(nbrFound[0])
			nbr2, _ = strconv.Atoi(nbrFound[1])
			instanceRes = nbr1 * nbr2
		}
		totalRes += instanceRes
	}

	fmt.Printf("Total result is %v\n", totalRes)
}
