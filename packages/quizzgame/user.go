package quizzgame

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/fatih/color"
)

type User struct {
	Pseudo string
	Score  int
}

func CreateUser() User {
	sc := bufio.NewScanner(os.Stdin)

	var user User

	for {
		fmt.Printf("Votre pseudo (<= 8 caractères): \n")
		color.Cyan("=> ")
		if sc.Scan() {
			user.Pseudo = sc.Text()
			lengthPseudo := utf8.RuneCountInString(user.Pseudo)
			if lengthPseudo > 8 {
				color.Red("Pseudo trop long, veuillez entrer un pseudo de 8 caractères ou moins.\n")
				continue
			} else {
				color.Magenta("Pseudo: %s\n", user.Pseudo)
				fmt.Println()
			}
		} else {
			color.Red("Error for the read user input:", sc.Err())
			return User{}
		}

		break
	}

	return user
}

func (u *User) SetScore(score int) {
	u.Score = score
}
