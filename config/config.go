// Package config will read and return config settings
package config

import (
    "fmt"
)

type Config struct {
    server string
    port string
}

func ReadConfig(filename string) (string, string, error) {

    fmt.Println("ReadConfig function call works")
    return "", "", nil
}
