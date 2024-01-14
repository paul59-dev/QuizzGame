package dashboard

import (
	"bufio"
	"fmt"
	"os"
	"paul/quizzGame/packages/quizzgame"
	"strconv"
)

func Dashboard() {
	sc := bufio.NewScanner(os.Stdin)

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
			number, err := strconv.Atoi(userInput)
			if err != nil {
				fmt.Println("Error to convert variable to interger: ", err)
				return
			}

			if number < 0 || number > len(dashboard) {
				fmt.Printf("Nombre invalide, veuillez entrer un nombre entre 0 et %d\n", len(dashboard))
				continue
			}

			switch number {
			case 1:
				quizzgame.QuizzGame()
			default:
				fmt.Println("Nombre invalide !")
			}
		} else {
			fmt.Println("Error for the read user input:", sc.Err())
			return
		}

		break
	}

	fmt.Println()
}
