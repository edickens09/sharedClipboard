package main

import (
	"fmt"

	"github.com/edickens09/sharedClipboard/commands"
        "github.com/edickens09/sharedClipboard/clipboard"
        "github.com/edickens09/sharedClipboard/config"
)

func main() {
        file := "clipboard.dat"
	copied := "Hard coded text"

        config.ReadConfig("config.yaml")

	commands.Copy(copied)
	
        if err := clipboard.CopyToClipboard(file); err != nil {
            fmt.Println("Error writing to file")
            fmt.Println(err)
        }
}	
