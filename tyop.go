package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	const (
		homeLower = "asdfghjkl;'"
	)

	var (
		amtLines   uint
		chunkSize  uint
		lineLength uint
		seed       int64

		line []byte
		i    uint
		k    uint
	)

	flag.UintVar(&amtLines, "n", 3, "amount of lines")
	flag.UintVar(&chunkSize, "c", 5, "size of chunks to break each line into")
	flag.UintVar(&lineLength, "l", 20, "length of each line, excluding spaces")
	flag.Int64Var(&seed, "s", time.Now().UnixNano(), "seed for RNG")
	flag.Parse()

	if amtLines == 0 {
		panic("Amount of lines must be positive")
	}
	if lineLength == 0 {
		panic("Line length must be positive")
	}

	line = make([]byte, lineLength)
	rand.Seed(seed)
	scanner := bufio.NewScanner(os.Stdin)

	for k = amtLines; k > 0; k-- {
		// Construct line.
		for i = 0; i < lineLength; i++ {
			line[i] = homeLower[rand.Intn(len(homeLower))]
		}

		// Print line.
		for i = 0; i < lineLength; i++ {
			if chunkSize > 0 && i%chunkSize == 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%c", line[i])
		}

		// Accept arbitrary input.
		fmt.Println()
		scanner.Scan()
	}

	fmt.Printf("Happy typing! (Seed: %v)", seed) // Print seed for replication.
}
