# README

Tyop is a simple program to generate strings for typing practice. It's currently implemented in Go 1.8.

## Building

Run `go fmt && go build`.

## Usage

Running `tyop` without flags seeds the RNG with the current time and generates 3 lines of 20 characters each separated at intervals of 5 characters.

### Flags

`-c` The size of chunks to break each line into (0 to disable chunking)

`-l` The length of each line

`-n` The amount of lines to generate

`-s` The seed value for the RNG


## Cleaning

Run `go clean`.
