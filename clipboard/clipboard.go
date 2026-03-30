//Package clipboard is for interacting with clipboard and file
package clipboard

import (
	"os"
        "encoding/csv"

        "github.com/edickens09/sharedClipboard/commands"
)

type ClipboardContent struct {
    
    content string
    version int64
}

func CreateClipboard() error {
    file, err := os.OpenFile("clipboard.dat", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644,)
    if err != nil {

        return err
    }
    defer file.Close()

    return nil
}

func CopyToClipboard() error {


    file, err := os.Open("clipboard.dat")
    if err != nil {
        return err
    }

    defer file.Close()
    
    copiedText := commands.Paste()
    if copiedText != 

    //copy from system clipboard to file make sure and increment the file

    return nil
    
}

func CopyFromClipboard( text string) error {
    file, err := os.Open("clipboard.dat")
    if err != nil {
        return err
    }

    defer file.Close()
    return nil
}
