package lib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ConsoleApp(prompt io.Reader) {
	fmt.Println("Masukan Perintah: (create|edit|delete|display)")

	command := bufio.NewScanner(prompt)
	console := NewApplication()

	for {
		command.Scan()
		readCmd := strings.ToLower(command.Text())
		RunCommand(readCmd, console)

	}
}

func RunCommand(command string, app *Application) {
	switch command {
	case "display":
		app.CommandDisplay()
	case "edit":
		msg, err := app.CommandEdit()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(msg)
		}

	case "delete":
		msg, err := app.CommandDelete()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(msg)
		}
	case "create":
		msg := app.CommandCreate()
		fmt.Println(msg)
	case ".exit":
		os.Exit(1)
	default:
		fmt.Println("perintah tidak dikenal")
	}
}
