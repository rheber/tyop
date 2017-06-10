# README

Tyop is a simple program to generate strings for typing practice. It's currently implemented in Go 1.8.

## Building

Run `go fmt && go build`.

## Usage

Running `tyop` without flags seeds the RNG with the current time and generates 3 lines of 20 home row characters each separated at intervals of 5 characters.

### Flags

`-c` The size of chunks to break each line into (0 to disable chunking)

`-l` The length of each line

`-n` The amount of lines to generate

`-r` A string of characters corresponding to sets of keys as follows:

* n `1234567890-=

* N ~!@#$%^&*()_+

* t qwertyuiop[]\

* T QWERTYUIOP{}|

* h asdfghjkl;'

* H ASDFGHJKL:"

* b zxcvbnm,./

* B ZXCVBNM<>?

`-s` The seed value for the RNG

`-x` A string of extra characters to practice

## Cleaning

Run `go clean`.
