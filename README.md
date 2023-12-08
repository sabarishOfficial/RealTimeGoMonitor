# Go Server Monitor

This GoLang project provides a simple and real-time monitoring solution for your server's RAM and CPU usage. The application updates the information every 2 seconds, providing a live view without the need for manual page refresh.

## Features

- **Real-time Monitoring:** The application continuously monitors your server's RAM and CPU usage, providing up-to-date information.

- **Automatic Updates:** Information is automatically updated every 2 seconds, ensuring you have a live view of your server's performance.

## Prerequisites

Make sure you have the following GoLang packages installed:

- [github.com/shirou/gopsutil/cpu](https://pkg.go.dev/github.com/shirou/gopsutil/cpu)
- [github.com/shirou/gopsutil/mem](https://pkg.go.dev/github.com/shirou/gopsutil/mem)

You can install them using:

```bash
go get -u github.com/shirou/gopsutil/cpu
go get -u github.com/shirou/gopsutil/mem
```
## To run the project, execute the following command in your terminal:
```bash
go run main.go
```
After running the command, open your browser and navigate to
```bash
http://localhost:8000/
```
# Usage
The monitoring interface provides a clear visualization of your server's RAM and CPU usage. With automatic updates every 2 seconds, you can keep a close eye on your system's performance.
# Contributing
We welcome contributions! If you'd like to contribute to this project, feel free to open issues, create pull requests, or provide feedback.