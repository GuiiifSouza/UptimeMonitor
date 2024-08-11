package checker

import (
    "fmt"
    "net"
    "net/http"
    "strings"
    "time"
    "UptimeMonitor/internal/utils"
)

// CheckWebsite checks the status of a website
func CheckWebsite(url string) string {
    protocol, host := extractProtocolAndHost(url)
    port := "80"
    if protocol == "https" {
        port = "443"
    }

    if !checkTCPConnection(host, port) {
        return utils.Red("DOWN")
    }

    client := http.Client{
        Timeout: time.Duration(utils.AppConfig.TimeoutSeconds) * time.Second, // Use the timeout from the config
    }

    resp, err := client.Get(url)
    if err != nil {
        if isTimeoutError(err) {
            return utils.Yellow("TIMEOUT")
        }
        return utils.Red("DOWN")
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        return utils.Green("UP")
    }

    return utils.Red(fmt.Sprintf("DOWN (Status Code: %d)", resp.StatusCode))
}

// Extract the protocol and host from the URL
func extractProtocolAndHost(url string) (protocol, host string) {
    if strings.HasPrefix(url, "https://") {
        return "https", strings.TrimPrefix(url, "https://")
    }
    if strings.HasPrefix(url, "http://") {
        return "http", strings.TrimPrefix(url, "http://")
    }
    return "", url
}

// Check the TCP connection to the host and port
func checkTCPConnection(host, port string) bool {
    timeout := time.Duration(utils.AppConfig.TimeoutSeconds) * time.Second
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
    if err != nil {
        return false
    }
    conn.Close()
    return true
}

// Check if the error is a timeout error
func isTimeoutError(err error) bool {
    netErr, ok := err.(net.Error)
    return ok && netErr.Timeout()
}

