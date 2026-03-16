package commands

import (
	"fmt"
	"os"

	cb "github.com/edickens09/go-clipboard/clipboard"
)

func Copy(text string) {

	c := cb.New()

	if err := c.CopyText(text); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
} 

func Paste() (text string) {
	c := cb.New()

	text, err := c.PasteText()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
