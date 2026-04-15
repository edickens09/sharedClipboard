package main

import (
	"fmt"

	"github.com/edickens09/sharedClipboard/commands"
        "github.com/edickens09/sharedClipboard/clipboard"
//	"go.yaml.in/yaml/v4"
)

type Config struct {
	server string
        port string
}

func main() {
	copied := "Hard coded text\n"

	commands.Copy(copied)
	
        if err := clipboard.CreateClipboard(); err != nil {
            fmt.Println("Error creating clipboard")
        }

        if err := clipboard.CopyToClipboard(); err != nil {
            fmt.Println("Error writing to file")
        }
}	
