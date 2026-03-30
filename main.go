package main

import (
	"fmt"

	"github.com/edickens09/sharedClipboard/commands"
//	"go.yaml.in/yaml/v4"
)

type Config struct {
	server string
}

func main() {
	copied := "Hard coded text\n"

	commands.Copy(copied)
	
	pasted := commands.Paste()
	fmt.Printf(pasted)
}	
