# Pre-Requisites
- Golang installed, currently using: go 1.14.11 windows/amd64 in Jetbrains Goland IDE
- A basic understanding of how to install Go, configure your GOPATH, etc

# Installation
1. Make sure this folder is somewhere under your GOPATH (environment variable)
2. Depending on your IDE, you may need to run `go get ./...` from the base folder

# Running the simulation
## Debugging / Running from source
Run this from the current directory 
- `go run main.go`
- `go run main.go -milliseconds 500`
- `go run main.go -path auctionItems.json`
- `go run main.go -milliseconds 500 -path auctionItems.json`


## Understanding the flags
- `-milliseconds`: used to determine the item process time. I.E. A value of 500 would indicate 500 milliseconds per item, meaning 10 items in the auction would finish in 5 seconds.
- `-path`: used to determine the location of the JSON file that will contain an array of auction items. See `auctionItems.json` file in this directory for an example.

## Running the executable

## Running with 