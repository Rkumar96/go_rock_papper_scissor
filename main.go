package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	playerScore   int
	computerScore int
)

func main() {
	fmt.Println(`
	███████████████████████████████████████████████████████████████████████████████████████████████████████
	█▄─▄▄▀█─▄▄─█─▄▄▄─█▄─█─▄███▄─▄▄─██▀▄─██▄─▄▄─█▄─▄▄─█▄─▄▄─█▄─▄▄▀███─▄▄▄▄█─▄▄▄─█▄─▄█─▄▄▄▄█─▄▄▄▄█─▄▄─█▄─▄▄▀█
	██─▄─▄█─██─█─███▀██─▄▀█████─▄▄▄██─▀─███─▄▄▄██─▄▄▄██─▄█▀██─▄─▄███▄▄▄▄─█─███▀██─██▄▄▄▄─█▄▄▄▄─█─██─██─▄─▄█
	▀▄▄▀▄▄▀▄▄▄▄▀▄▄▄▄▄▀▄▄▀▄▄▀▀▀▄▄▄▀▀▀▄▄▀▄▄▀▄▄▄▀▀▀▄▄▄▀▀▀▄▄▄▄▄▀▄▄▀▄▄▀▀▀▄▄▄▄▄▀▄▄▄▄▄▀▄▄▄▀▄▄▄▄▄▀▄▄▄▄▄▀▄▄▄▄▀▄▄▀▄▄▀`)

	fmt.Println("Enter your Name")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)

	fmt.Printf("%s %s\n", "welcome, Do you want to start Playing", name)

	for {
		fmt.Println(`Press "s" to start `)
		fmt.Println(`Press "e" to exit `)
		fmt.Println(`Press "v" to View score `)

		//reading single character from user
		reader := bufio.NewReader(os.Stdin)
		char, _, error := reader.ReadRune()

		if error != nil {
			fmt.Println("Error Occured")
			fmt.Println(error)
			break
		}

		switch char {
		case 's':
			fmt.Println("Game is Starting")
			startGame()
			printScore(name)
		case 'e':
			fmt.Println("Exiting ...")
			os.Exit(1)
		case 'v':
			printScore(name)
		default:
			fmt.Println("Wrong input Please Try Again later")

		}
	}

}

func printScore(name string) {
	fmt.Printf("\n%s %d : %d %s\n\n", name, playerScore, computerScore, "Computer")
}

func startGame() {
	fmt.Println(`Press "r"  for Rock 🥌`)
	fmt.Println(`Press "p"  for Paper 📰`)
	fmt.Println(`Press "s"  for Scissor ✂️`)

	reader := bufio.NewReader(os.Stdin)
	userChoice, _, error := reader.ReadRune()

	if error != nil {
		fmt.Println(error)
	}

	computerChoice := getComputerChoice()

	// fmt.Println("You choose " + string(userChoice))
	choices := map[string]string{"r": "Rock", "p": "Paper", "s": "Scissor"}

	fmt.Println("Computer chooses " + choices[string(computerChoice)])

	result(userChoice, computerChoice)

}

func getComputerChoice() rune {
	choices := [3]rune{'r', 'p', 's'}
	return choices[rand.Intn(len(choices))]
}

func result(user rune, computer rune) {

	combinedResults := string(user) + string(computer)
	switch combinedResults {
	case "rs", "pr", "sp":
		fmt.Println("You Won 🏆")
		playerScore++
	case "rp", "ps", "sr":
		fmt.Println("You Lost 💩")
		computerScore++
	case "rr", "pp", "ss":
		fmt.Println("Match Draws 🤙🏻")
	}
}
