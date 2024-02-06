package main

import (
	"bufio"
	"os"

	"crud-cli/lib"
)

func main() {
	lib.ConsoleApp(bufio.NewReader(os.Stdin))
}