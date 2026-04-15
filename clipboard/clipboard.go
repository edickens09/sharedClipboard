//Package clipboard is for interacting with clipboard and file
package clipboard

import (
	"os"
        "io"
        "encoding/csv"
        "strconv"

        "github.com/edickens09/sharedClipboard/commands"
)

type ClipboardContent struct {
    
    content string
    version int64
}

var cbContent ClipboardContent

func CreateClipboard() error {
    file, err := os.OpenFile("clipboard.dat", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644,)
    if err != nil {

        return err
    }
    defer file.Close()

    return nil
}

func CopyToClipboard() error {

    var csvContent []string

    file, err := os.Open("clipboard.dat")
    if err != nil {
        return err
    }

    defer file.Close()
    
    copiedText := commands.Paste()
    
    reader := csv.NewReader(file)

    content, err := reader.Read()
    if err != nil {
        if err == io.EOF {

        } else {
            return err
        }
    }

    if copiedText != content[0] {
        writer := csv.NewWriter(file)
        
        newVersionInt, err := strconv.ParseInt(content[1], 10, 64)
        if err != nil {
            return err
        }

        newVersionInt += 1
        newVersion := strconv.FormatInt(newVersionInt, 10)

        csvContent = append(csvContent, copiedText)
        csvContent = append(csvContent, newVersion)

        if err = writer.Write(csvContent); err != nil {
            return err
        }

        writer.Flush()
    }

    return nil
    
}

func CopyFromClipboard( text string) error {

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
    cbContent.version, err = strconv.ParseInt(content[1], 10, 64)
    if err != nil {
        return err
    }

    commands.Copy(cbContent.content)
    
    return nil
}
