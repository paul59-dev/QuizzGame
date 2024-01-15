package quizzgame

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"paul/quizzGame/packages/colors"
	"strings"
	"unicode"

	"github.com/fatih/color"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func QuizzGame() {
	sc := bufio.NewScanner(os.Stdin)

	var (
		userInput string // => ""
		score     int    // => 0
	)

	// Create User
	user := CreateUser()

	// Pointer
	scorePtr := &score

	// Open problem.csv in read mode
	readCSV, err := os.Open("problem.csv")
	if err != nil {
		color.Red("imposible open the csv file: ", err)
	}
	defer readCSV.Close()

	reader := csv.NewReader(readCSV)

	// Read all data in the csv file
	lines, err := reader.ReadAll()
	if err != nil {
		color.Red("Error of the read csv file: ", err)
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

	// Given response with differente question of the csv file
	// Principal loop
	for i := 0; i < len(questions); i++ {
		fmt.Printf("Question %d: %s:\n", i+1, questions[i])
		colors.Color("Réponse: ", "cyan")

		// User input response
		if sc.Scan() {
			// Delete accent, convert to lower case
			userInput = removeAccents(strings.ToLower(sc.Text()))
			responses[i] = removeAccents(strings.ToLower(responses[i]))
			if userInput == responses[i] {
				colors.Color("Bravo, continuer !", "green")
				*scorePtr++ // Increment score for memory adress of the score variable
				fmt.Println()
			} else {
				colors.Color("Désolé, réessayer !", "yellow")
				fmt.Println()
				colors.Color("Votre score: %#v", "cyan", *scorePtr)
				fmt.Println()
				break
			}
		} else {
			color.Red("Error for the read user input:", sc.Err())
			return
		}
	}

	user.Score = *scorePtr // Inject score pointer to attribut Score in the User struct

	// Open or create user.csv file
	writerCSV, err := os.OpenFile("user.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		color.Red("Impossible to open user.csv: ", err)
		return
	}
	defer writerCSV.Close()

	// New writer in user.csv
	writer := csv.NewWriter(writerCSV)
	defer writer.Flush()

	// Inject the data to user.csv
	err = writer.Write([]string{user.Pseudo, fmt.Sprint(user.Score)})
	if err != nil {
		color.Red("Error writting to csv: ", err)
		return
	}

	colors.Color("\nDonnées de l'utilisateur enregistrées dans user.csv.", "green")
}

// Function allow delete the accent in the word
func removeAccents(input string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Retirer les marques diacritiques
	}))
	result, _, _ := transform.String(t, input)
	return result
}
