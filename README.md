# Uptime Monitor

## Overview

Uptime Monitor is a Go application designed to check the availability of websites. It monitors a list of URLs and prints their status to the terminal. The status is updated continuously, and the display is color-coded to indicate whether a website is "UP" or "DOWN".

## Features

- **Configurable Monitoring**: URLs and settings are defined in a JSON configuration file.
- **Real-Time Status Updates**: Monitors and updates the status of websites every few seconds.
- **Color-Coded Output**: Displays website status with color codes for easy identification.

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/GuiiifSouza/UptimeMonitor.git
    cd UptimeMonitor
    ```

2. **Install dependencies:**

    Ensure you have Go installed and run:

    ```bash
    go mod download
    ```

3. **Run the application:**

    ```bash
    go run cmd/main.go
    ```

## Configuration

The application reads its configuration from `config.json`. This file should include:

- **URLs**: A list of URLs to monitor.
- **Timeout**: The timeout for each request (in seconds).
- **Interval**: The time between status checks (in seconds).

### Example `config.json`

```json
{
  "urls": [
    "https://www.google.com",
    "https://www.github.com"
  ],
  "timeout": 2,
  "interval": 2
}
```

## Directory Structure

- `cmd/main.go`: Entry point for the application.
- `config.json`: Configuration file with URLs and settings.
- `internal/checker/checker.go`: Contains the logic to check the availability of websites.
- `internal/utils/fileloader.go`: Handles loading URLs and configuration.
- `internal/utils/utils.go`: Utility functions including color codes and screen clearing.

## Logging

The application prints the status of each URL directly to the terminal. Logging is handled within the application and no external log files are used.

## Contributing

Feel free to fork the repository and submit pull requests with improvements. Ensure that your contributions are well-documented and include relevant tests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

