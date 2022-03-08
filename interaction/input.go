package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const PlayerNameLength = 10

var reader = bufio.NewReader(os.Stdin)

func GetPlayerName() string {
	for {
		userName, _ := getPlayerInput("Howdy! Enter your name please: ")

		if userName == "" {
			fmt.Println("Name can't be blank! Please try again.")
		} else if len(userName) > PlayerNameLength {
			textErr := fmt.Sprintf("Name must be lower than %v characters. Please try again.", PlayerNameLength)
			fmt.Println(textErr)
		} else {
			return userName
		}
	}
}

func getPlayerInput(promptText string) (string, error) {
	fmt.Print(promptText)

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput, nil
}