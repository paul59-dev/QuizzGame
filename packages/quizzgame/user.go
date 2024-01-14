package quizzgame

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Pseudo string
	Score  int
}

func CreateUser() User {
	sc := bufio.NewScanner(os.Stdin)

	var user User

	fmt.Printf("Votre pseudo (<= 10 caractères): ")
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
		return User{}
	}

	return user
}

func (u *User) SetScore(score int) {
	u.Score = score
}
