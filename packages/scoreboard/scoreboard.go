package scoreboard

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Scoreboard() {
	// Open problem.csv in read mode
	readCSV, err := os.Open("user.csv")
	if err != nil {
		if os.IsNotExist(err) {
			color.Yellow("Aucune donnée pour le moment")
		} else {
			color.Red("imposible open the csv file: ", err)
		}
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
	pseudos := make([]string, len(lines))
	for i, line := range lines {
		if len(line) > 1 {
			pseudos[i] = line[0]
		}
	}

	// Read second colomn in the csv file
	scores := make([]string, len(lines))
	for i, line := range lines {
		if len(line) > 1 {
			scores[i] = line[1]
		}
	}

	// Ajouter des espaces pour les étiquettes "Pseudo" et "Score"
	maxPseudoWidth := maxLength(pseudos) + 2 // +2 pour compenser les espaces ajoutés
	//maxScoreWidth := maxLength(scores) + 2   // +2 pour compenser les espaces ajoutés

	// Afficher les données sous forme de tableau
	fmt.Println("+------------+----------+")
	fmt.Println("| Pseudo     |  Score   |")
	fmt.Println("+------------+----------+")

	for i := 0; i < len(scores); i++ {
		fmt.Printf("| %-*s | %-6s |\n", maxPseudoWidth, pseudos[i], scores[i])
	}

	fmt.Println("+------------+----------+")
}

func maxLength(strings []string) int {
	max := 0
	for _, s := range strings {
		if len(s) > max {
			max = len(s)
		}
	}
	return max
}
