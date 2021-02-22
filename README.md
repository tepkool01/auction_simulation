# Overview and Foreword
This repository is a homework assignment given by Lookout which implements an Auction based on a JSON file. 

The entry point is in `main.go` (in this top-level folder). Other applications can leverage this Auction system by extracting the auction folder which contains the auction package. 

This assignment was time-boxed at 48 hours. Of those hours, I had an 8-hour school day on Saturday, and general life maintenance. 
I'm also not an expert at GoLang, but don't think that will get in the way of showing my approach to problems/creative ways to bastardize variable names.

## TL;DR
Go to the `artifacts` folder, and grab either `lookout_interview_linux` or `lookout_interview_windows.exe` (depending on your OS, and I hope you are on 64 bit).
And run it from the command line. Check the docs later for explanation on how to use it with flags. 

## Folder explanation
- `/auction` is where the core package is located. This includes the models, interfaces, tests, and test data (yes, golang likes their tests adjacent to the code)
  - for the most part, the interesting concrete 'class' is `auction.go`
- `/auction/model` where the models (bids, item listings, type enums) is located
- `/artifacts` is the location of a linux and windows binary to run this (additional OSs on request)

## Pre-Requisites
- Golang installed, currently using: **go 1.14.11** windows/amd64 in Jetbrains Goland IDE
- A basic understanding of how to install Go, configure your GOPATH, etc is useful

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

# Architecture Diagrams
A UML diagram of how the various structs (class-like objects of sorts) are connected with one another. 

![Diagram of System](artifacts/uml_architecture.png?raw=true "Architecture")

# Notes / Future Design
The assignment mentioned a 'production-ready', and I think that depends on the definition of who it is serving. I would not expose this externally as
an actual auction system. I think internal clients would be fine, followed by iterating on the design. Other things I've considered but prioritized lower (and did not get to) are as follows:
1. Security -- this is a security company after all. Funny that I placed this low on my list. Particularly around the opening of files / general attack surfaces of input validation. 
2. Concurrency -- there's a job-worker pattern I use all the time that I put into `artifacts/workerpool.go`, but there was some confusion on my end at the concurrency requirement for this assignment.
3. Usability -- better documentation, better error messages to help guide users
4. Logging -- Debug logs around the system to diagnose issues 
5. Testing -- Improve fail cases, improve coverage, particularly branch coverage
   - Possibly system/stress testing if multiple files were introduced
6. Structure -- Improve folder structure to scale out the system better
7. [Possibly] a separate class to handle all text and saving that into a template file
8. A sequence diagram in the architectural artifacts showing the program flow (happy path)
