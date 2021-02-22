# Pre-Requisites
- Golang installed, currently using: <b>go 1.14.11</b> windows/amd64 in Jetbrains Goland IDE
- A basic understanding of how to install Go, configure your GOPATH, etc is useful

# Overview and Foreword
This repository is a homework assignment given by Lookout which implements an Auction based on a JSON file. 

The entry point is in `main.go` (in this top-level folder). Other applications can leverage this Auction system by extracting the auction folder which contains the auction package. 

This assignment was time-boxed at 48 hours. Of those hours, I had an 8-hour school day on Saturday, and general life maintenance. 
I'm also not an expert at GoLang, but don't think that will get in the way of showing my approach to problems/creative ways to bastardize variable names.

## Installation
1. Make sure this folder is somewhere under your GOPATH (environment variable)
2. Depending on your IDE, you may need to run `go get ./...` from the base folder

## Running the simulation
The entry point, main.go, allows users to test the functionality of the auction package. Below will detail the flags, build commands, etc.

## Understanding the flags
- `-milliseconds`: used to determine the item process time. I.E. A value of 500 would indicate 500 milliseconds per item, meaning 10 items in the auction would finish in 5 seconds.
- `-path`: used to determine the location of the JSON file that will contain an array of auction items. See `auctionItems.json` file in this directory for an example.

## Debugging / Running from source
Run this from the current directory 
- `go run main.go`
- `go run main.go -milliseconds 500`
- `go run main.go -path auctionItems.json`
- `go run main.go -milliseconds 500 -path auctionItems.json`
- <b>Testing Command</b>, (switch to auction and run the test) `cd auction && go test`

## Building the binary
- `go build` and this should create a binary based on your system (windows, linux, mac, etc)

## Running the executable
Running the executable is the same as the debugging steps above. Use the flags as needed (defaults are provided). The exception is that instead of 
`go run main.go`, you would simply execute the file, i.e. `lookout_interview.exe -path auctionItems.json` for windows.

# Notes / Future Design
The assignment mentioned a 'production-ready', and I think that depends on the definition of who it is serving. I would not expose this externally as
an actual auction system. I think internal clients would be fine, followed by iterating on the design. Other things I've considered but prioritized lower (and did not get to) are as follows:
1. Security -- this is a security company after all. Funny that I placed this low on my list. Particularly around the opening of files / general attack surfaces of input validation. 
2. Concurrency -- there's a job-worker pattern
3. Usability -- better documentation, better error messages to help guide users
4. Logging -- Debug logs around the system to diagnose issues 
5. Testing -- Improve fail cases, improve coverage, particularly branch coverage
6. Structure -- Improve folder structure to scale out the system better