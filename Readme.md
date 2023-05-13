# Go Script Runner

Go Script Runner is a tool that runs scripts from a folder, logs the response to a log file, and sends a message to a Discord webhook. It provides a convenient way to automate the execution of scripts and keep track of their outputs.

## Features

- Executes scripts from a folder
- Logs the script responses to a log file
- Sends a message to a Discord webhook
- Supports setting a cron schedule for periodic execution

## Prerequisites

Before using the Go Script Runner, ensure that you have the following:

- Go programming language installed (version 1.20 or higher)
- A folder named "scripts" in the root directory, containing the scripts you want to execute
- Set the `DISCORD_WEBHOOK_URL` environment variable with your Discord webhook URL
- Set the `CRON_SCHEDULE` environment variable to define the schedule for periodic execution

All environment variables can be set via a .env file.

## Logging

The tool logs the script responses to a log file named `scripts.log`. The log file is created in the same directory as the tool. Each script execution is appended to the log file along with a timestamp.

## Installation

To install and use the Go Script Runner, follow these steps:

### 1. Clone the repository

```sh
git clone https://github.com/dontWatchMeCode/pipe
```

### 2. Build the project

```sh
make build
```

you can also build the binary via a docker container

```sh
make docker-build
```

currently only linux is supported

### 3. Run the tool

```sh
./main
```
