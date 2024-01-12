package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	var userInput string

	file, err := os.Open("problem.csv")
	if err != nil {
		fmt.Errorf("imposible open the csv file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all data in the csv file
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error of the read csv file", err)
		return
	}

	// Read first colomn in the csv file
	questions := make([]string, len(lines))
	for i, line := range lines {
		if len(line) > 1 {
			questions[i] = line[0]
		}
	}

	responses := make([]string, len(lines))
	for i, line := range lines {
		if len(line) > 1 {
			responses[i] = line[1]
		}
	}

	// Given response with differente question of the csv file
	for i := 0; i < len(questions); i++ {
		fmt.Printf("Question %d: %s:\n", i+1, questions[i])
		fmt.Print("Response: ")

		// Ueser input
		sc := bufio.NewScanner(os.Stdin)
		if sc.Scan() {
			userInput = sc.Text()
			fmt.Printf("Value of input: %s\n", userInput)
			if userInput == responses[i] {
				fmt.Println("Yes, continious !")
			} else {
				fmt.Println("Sorry, restarting !")
				break
			}
		} else {
			fmt.Println("Error for the read user input:", sc.Err())
			return
		}
	}
}
