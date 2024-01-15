package quizzgame

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

type User struct {
	Pseudo string
	Score  int
}

func CreateUser() User {
	sc := bufio.NewScanner(os.Stdin)

	var user User

	for {
		fmt.Printf("Votre pseudo (<= 8 caractères): ")
		fmt.Print("=> ")
		if sc.Scan() {
			user.Pseudo = sc.Text()
			lengthPseudo := utf8.RuneCountInString(user.Pseudo)
			if lengthPseudo > 8 {
				fmt.Printf("Pseudo trop long, veuillez entrer un pseudo de 8 caractères ou moins.\n")
				continue
			} else {
				fmt.Printf("Pseudo: %s\n", user.Pseudo)
				fmt.Println()
			}
		} else {
			fmt.Println("Error for the read user input:", sc.Err())
			return User{}
		}

		break
	}

	return user
}

func (u *User) SetScore(score int) {
	u.Score = score
}
