package main

import (
    "fmt"
    "time"
    "UptimeMonitor/internal/checker"
    "UptimeMonitor/internal/utils"
)

func main() {
    // Load configurations
    err := utils.LoadConfig("config.json")
    if err != nil {
        fmt.Println("Error loading configuration:", err)
        return
    }

    for {
        utils.ClearScreen()
        fmt.Println("Uptime Monitor")
        fmt.Println("===================")
        for _, url := range utils.AppConfig.Sites {
            status := checker.CheckWebsite(url)
            fmt.Printf("%s : %s\n", url, status)
        }
        time.Sleep(time.Duration(utils.AppConfig.IntervalSeconds) * time.Second)
    }
}

