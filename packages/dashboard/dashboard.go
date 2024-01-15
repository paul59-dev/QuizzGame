package dashboard

import (
	"bufio"
	"fmt"
	"os"
	"paul/quizzGame/packages/colors"
	"paul/quizzGame/packages/quizzgame"
	"paul/quizzGame/packages/scoreboard"
	"strconv"

	"github.com/fatih/color"
)

func Dashboard() {
	sc := bufio.NewScanner(os.Stdin)

	var dashboard = []string{
		"QuizzGame",
		"Voir les scores",
	}

	fmt.Println()
	colors.Color("Menu principal !\n", "yellow")
	for index, value := range dashboard {
		fmt.Printf("   %d - %s\n", index+1, value)
	}

	for {
		colors.Color("Enter number: ", "blue")
		if sc.Scan() {
			userInput := sc.Text()
			number, err := strconv.Atoi(userInput)
			if err != nil {
				color.Red("Error to convert variable to interger: ", err)
				return
			}

			if number < 0 || number > len(dashboard) {
				color.Red("Nombre invalide, veuillez entrer un nombre entre 0 et %d\n", len(dashboard))
				continue
			}

			switch number {
			case 1:
				quizzgame.QuizzGame()
			case 2:
				scoreboard.Scoreboard()
			default:
				color.Red("Nombre invalide !")
			}
		} else {
			color.Red("Error for the read user input:", sc.Err())
			return
		}

		break
	}

	fmt.Println()
}
