//Package clipboard is for interacting with clipboard and file
package clipboard

import (

	"os"
        "io"
        "fmt"
        "encoding/csv"
        "strconv"

        "github.com/edickens09/sharedClipboard/commands"
        "github.com/fsnotify/fsnotify"
)

type ClipboardContent struct {
    
    content string
    version int64
}

var cbContent ClipboardContent

func CreateClipboard() error {
    file, err := os.OpenFile("clipboard.dat", os.O_CREATE, 0644,)
    if err != nil {

        return err
    }
    defer file.Close()

    return nil
}

func CopyToClipboard(filename string) error {
    var version int64
    
    //find a way to open file and overwrite it each time
    file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        return err
    }

    defer file.Close()
    
    //getting the clipboard info into memory
    copiedText := commands.Paste()
    
    reader := csv.NewReader(file)

    content, err := reader.Read()
    if err != nil {
        if err == io.EOF {
            //should something else be in this area?

        } else {
            return err
        }
    }

    if len(content) == 0 {

        version = 0

        if err = WriteToCSV(file, version, copiedText); err != nil {
            return err
        }

        return nil

    }

    if copiedText != content[0] {

        version, err = strconv.ParseInt(content[1], 10, 64)
        if err != nil {
            return err
        }
        
        if err = WriteToCSV(file, version, copiedText); err != nil {
            return err
        }
    }

    return nil
}

func CopyFromClipboard(text string, filename string) error {
    file, err := os.Open(filename)
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

func WriteToCSV(file *os.File, version int64, copiedText string) error {
    var csvContent []string

    _, err := file.Seek(0,0)
    if err != nil {
        return err
    }

    err = file.Truncate(0)
    if err != nil {
        return err
    }

    writer := csv.NewWriter(file)

    version += 1
    newVersion := strconv.FormatInt(version, 10)
    
    csvContent = append(csvContent, copiedText)
    csvContent = append(csvContent, newVersion)

    if err = writer.Write(csvContent); err != nil {
        return err
    }

    writer.Flush()

    return nil
}

func WatchFile(filename string) error {

    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return err
    }

    defer watcher.Close()

    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                fmt.Println("event:", event)
                if event.Has(fsnotify.Write) {
                    fmt.Println("modified file:", event.Name)
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                fmt.Println("error:", err)
            }
        }
    }()
    return nil
}
