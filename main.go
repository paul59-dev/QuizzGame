package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var sc *bufio.Scanner

type User struct {
	Pseudo string
	Score  int
}

func initScanner() {
	// Initialisation du scanner avec l'entrée standard (os.Stdin)
	sc = bufio.NewScanner(os.Stdin)
}

func dashboard() {
	initScanner()

	var dashboard = []string{
		"QuizzGame",
		"Voir les scores",
	}

	fmt.Println("Menu principal !")
	for index, value := range dashboard {
		fmt.Printf("%d - %s\n", index+1, value)
	}

	for {
		fmt.Print("Enter number: ")
		if sc.Scan() {
			userInput := sc.Text()
			userNumber, err := strconv.Atoi(userInput)
			if err != nil {
				fmt.Println("Error to convert variable to interger: ", err)
				return
			}

			if userNumber < 0 || userNumber > len(dashboard) {
				fmt.Printf("Nombre invalide, veuillez entrer un nombre entre 0 et %d", len(dashboard))
				continue
			}
		} else {
			fmt.Println("Error for the read user input:", sc.Err())
			return
		}

		break
	}

	fmt.Println()
}

// Function allow delete the accent in the word
func removeAccents(input string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Retirer les marques diacritiques
	}))
	result, _, _ := transform.String(t, input)
	return result
}

func main() {
	initScanner()

	var (
		userInput string // => ""
		score     int    // => 0
		user      User
	)

	// Pointer
	scorePtr := &score

	// Open problem.csv in read mode
	readCSV, err := os.Open("problem.csv")
	if err != nil {
		fmt.Println("imposible open the csv file: ", err)
	}
	defer readCSV.Close()

	reader := csv.NewReader(readCSV)

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

	// Create User
	fmt.Printf("Votre pseudo (<= 10 characteres): ")
	fmt.Print("=> ")
	if sc.Scan() {
		user.Pseudo = sc.Text()
		if len(user.Pseudo) > 10 {
			fmt.Printf("Pseudo trop long, veuillez entrer un pseudo de 10 caractères ou moins.\n")
		} else {
			fmt.Printf("Pseudo: %s\n", user.Pseudo)
			fmt.Println()
		}
	} else {
		fmt.Println("Error for the read user input:", sc.Err())
		return
	}

	dashboard()

	// Given response with differente question of the csv file
	// Principal loop
	for i := 0; i < len(questions); i++ {
		fmt.Printf("Question %d: %s:\n", i+1, questions[i])
		fmt.Print("Réponse: ")

		// User input response
		if sc.Scan() {
			// Delete accent, convert to lower case
			userInput = removeAccents(strings.ToLower(sc.Text()))
			responses[i] = removeAccents(strings.ToLower(responses[i]))
			if userInput == responses[i] {
				fmt.Println("Bravo, continuer !")
				*scorePtr++ // Increment score for memory adress of the score variable
				fmt.Println()
			} else {
				fmt.Println("Désolé, réessayer !")
				fmt.Println()
				fmt.Printf("Votre score: %#v", *scorePtr)
				fmt.Println()
				break
			}
		} else {
			fmt.Println("Error for the read user input:", sc.Err())
			return
		}
	}

	user.Score = *scorePtr // Inject score pointer to attribut Score in the User struct

	// Open or create user.csv file
	writerCSV, err := os.OpenFile("user.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Impossible to open user.csv: ", err)
		return
	}
	defer writerCSV.Close()

	// New writer in user.csv
	writer := csv.NewWriter(writerCSV)
	defer writer.Flush()

	// Inject the data to user.csv
	err = writer.Write([]string{user.Pseudo, fmt.Sprint(user.Score)})
	if err != nil {
		fmt.Println("Error writting to csv: ", err)
		return
	}

	fmt.Println("\nDonnées de l'utilisateur enregistrées dans user.csv.")
}
