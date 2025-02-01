# Task CLI

## About
A Task Tracker CLI made using Go for [roadmap.sh backend project](https://roadmap.sh/projects/task-tracker).

## Pre-requisite
Go 1.23.0

## Setup to Run Locally
1. Clone this Github repository
2. Go to the folder
3. Build executable file
```
go build cmd/task-cli.go
```

## Usage
### Commands
```
# Adding a new task
./task-cli add "<Task Description>"
# Output: Task added successfully (ID:1)

# Updating tasks
./task-cli update <ID> "<Task Description>"

# Deleting tasks
./task-cli delete <ID>

# Marking tasks as in-progress or done
./task-cli mark-in-progress <ID>
./task-cli mark-done <ID>

# Listing all tasks
./task-cli list

# Listing tasks by status
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```

### Usage Example
![task-cli](https://github.com/user-attachments/assets/f2868873-8d04-4be8-9a43-86b2cca4e7aa)
