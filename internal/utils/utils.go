package utils

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// AppConfig holds the configuration values
var AppConfig Config

// Config holds the configuration settings
type Config struct {
    Sites             []string `json:"sites"`
    TimeoutSeconds    int      `json:"timeout_seconds"`
    IntervalSeconds   int      `json:"interval_seconds"`
}

// LoadConfig loads the configuration from a JSON file
func LoadConfig(filepath string) error {
    file, err := os.Open(filepath)
    if err != nil {
        return err
    }
    defer file.Close()

    data, err := ioutil.ReadAll(file)
    if err != nil {
        return err
    }

    return json.Unmarshal(data, &AppConfig)
}

// ClearScreen clears the terminal screen
func ClearScreen() {
    fmt.Print("\033[H\033[2J")
}

// Color functions
func Red(text string) string {
    return "\033[31m" + text + "\033[0m"
}

func Green(text string) string {
    return "\033[32m" + text + "\033[0m"
}

func Yellow(text string) string {
    return "\033[33m" + text + "\033[0m"
}

