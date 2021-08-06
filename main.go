package main

import (
	"bufio"
	"fmt"
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
		case 'e':
			fmt.Println("Exiting ...")
			os.Exit(1)
		case 'v':
			printScore(name)

		}
	}

}

func printScore(name string) {
	fmt.Printf("\n%s %d : %d %s\n\n", name, playerScore, computerScore, "Computer")
}
