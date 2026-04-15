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
        file := "clipboard.dat"
	copied := "Hard coded text"

	commands.Copy(copied)
	
        if err := clipboard.CopyToClipboard(file); err != nil {
            fmt.Println("Error writing to file")
            fmt.Println(err)
        }
}	
