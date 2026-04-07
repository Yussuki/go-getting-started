package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {

	userTitle := getUserInput("Title: ")
	userContent := getUserInput("Content: ")

	userNote, err := note.New(userTitle, userContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	if err = userNote.Save(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Note saved sucessfully")
}

func getUserInput(promptText string) string {
	fmt.Print(promptText)

	// fmt.Scan and fmt.Scanln accepts only one value (one word or number, don't work with phrases)
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // Obs: This needs to be single quotes, because the parameter is a byte, not a char
	if err != nil {
		return ""
	}

	// Good practice, the ReadString includes the last "\n" and "\r" is included sometimes depending on OS
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
