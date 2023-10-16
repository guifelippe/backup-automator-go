# Backup Automator Go

Backup Automator Go is a Go-based tool that automates the task of backing up files and directories, generating zip files at specified intervals. This helps keep your data safe and allows for easy recovery on your schedule.

## Features

- Scheduled automatic backups.
- Choice of source files and directories.
- Compression in zip format.
- Flexible backup scheduling.

## Requirements

- Go 1.16 or higher.
- Write permission to the destination directory.

## Installation

Follow these steps to install and set up Backup Automator Go:

1. Clone the repository:

   ```shell
   git clone https://github.com/your-username/backup-automator-go.git

2. Navigate to the project directory:

    ```shell 
    cd backup-automator-go

3. Build the program:

    ```shell 
    go build

4. Run the program:

    ```shell 
    ./backup-automator-go

The program will start automatically running backups based on the configuration specified in the config.json file.