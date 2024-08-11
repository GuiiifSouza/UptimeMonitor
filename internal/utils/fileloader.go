package utils

import (
    "bufio"
    "os"
    "strings"
)

// LoadURLs loads the list of URLs from a file
func LoadURLs(filepath string) ([]string, error) {
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var urls []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" {
            urls = append(urls, line)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return urls, nil
}

