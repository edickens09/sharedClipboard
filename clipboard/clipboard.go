//Package clipboard is for interacting with clipboard and file
package clipboard

import (
	"os"
        "io"
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
/*
    if copiedText != content from csv {
        then write to file

        version += 1
    }
    copy from system clipboard to file make sure and increment the file
*/
    return nil
    
}

func CopyFromClipboard( text string) error {
    
    var cbContent ClipboardContent

    file, err := os.Open("clipboard.dat")
    if err != nil {
        return err
    }

    defer file.Close()
    
    reader := csv.NewReader(file)

    content, err := reader.Read()
    if err != nil {
        if err == io.EOF {
           //does something need to go here. I'm just trying to catch this value and move past it. 
        } else {
            return err
        }
    }

    cbContent.content = content[0]
    //leaving this asis for now, I would like version to remain an int but csv needs to be string?
    cbContent.version = content[1]
    
    return nil
}
