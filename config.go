package utils

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

type Config struct {
    TimeoutSeconds  int      `json:"timeout_seconds"`
    IntervalSeconds int      `json:"interval_seconds"`
    Sites           []string `json:"sites"`
}

var AppConfig Config

// LoadConfig loads configuration from a JSON file
func LoadConfig(filepath string) error {
    file, err := ioutil.ReadFile(filepath)
    if err != nil {
        return err
    }

    err = json.Unmarshal(file, &AppConfig)
    if err != nil {
        return err
    }

    return nil
}

