package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type User struct {
	Pseudo string
	Score  int
}

func removeAccents(input string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Retirer les marques diacritiques
	}))
	result, _, _ := transform.String(t, input)
	return result
}

func main() {

	var (
		userInput string // => ""
		score     int    // => 0
		user      User
	)

	// Pointer
	scorePtr := &score

	file, err := os.Open("problem.csv")
	if err != nil {
		fmt.Println("imposible open the csv file: ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all data in the csv file
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error of the read csv file: ", err)
		return
	}

	// Read first colomn in the csv file
	questions := make([]string, len(lines))
	for i, line := range lines {
		if len(line) > 1 {
			questions[i] = line[0]
		}
	}

	// Read second colomn in the csv file
	responses := make([]string, len(lines))
	for i, line := range lines {
		if len(line) > 1 {
			responses[i] = line[1]
		}
	}

	sc := bufio.NewScanner(os.Stdin)

	// Create User
	fmt.Printf("Votre pseudo (<= 10 characteres): ")
	fmt.Print("=> ")
	if sc.Scan() {
		user.Pseudo = sc.Text()
		if len(user.Pseudo) > 10 {
			fmt.Printf("Pseudo trop long, veuillez entrer un pseudo de 10 caractères ou moins.\n")
		} else {
			fmt.Printf("Pseudo: %s\n", user.Pseudo)
		}
	} else {
		fmt.Println("Error for the read user input:", sc.Err())
		return
	}

	// Given response with differente question of the csv file
	// Principal loop
	for i := 0; i < len(questions); i++ {
		fmt.Printf("Question %d: %s:\n", i+1, questions[i])
		fmt.Print("Response: ")

		// User input response
		if sc.Scan() {
			userInput = removeAccents(strings.ToLower(sc.Text()))
			responses[i] = removeAccents(strings.ToLower(responses[i]))
			fmt.Printf("Value of input: %s\n", userInput)
			fmt.Printf("Value of response: %s\n", responses[i])
			if userInput == responses[i] {
				fmt.Println("Yes, continious !")
				*scorePtr++
			} else {
				fmt.Println("Sorry, restarting !")
				fmt.Printf("Tour score is: %#v", *scorePtr)
				break
			}
		} else {
			fmt.Println("Error for the read user input:", sc.Err())
			return
		}
	}
}
