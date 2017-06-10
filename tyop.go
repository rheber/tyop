package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	const (
		numLower  = "`1234567890-="
		numUpper  = "~!@#$%^&*()_+"
		topLower  = "qwertyuiop[]\\"
		topUpper  = "QWERTYUIOP{}|"
		homeLower = "asdfghjkl;'"
		homeUpper = "ASDFGHJKL:\""
		botLower  = "zxcvbnm,./"
		botUpper  = "ZXCVBNM<>?"
	)

	var (
		amtLines   uint
		chunkSize  uint
		lineLength uint
		rowCode    string
		xtraChars  string
		seed       int64

		keys []byte
		line []byte
		i    uint
		k    uint
	)

	flag.UintVar(&amtLines, "n", 3, "amount of lines")
	flag.UintVar(&chunkSize, "c", 5, "size of chunks to break each line into")
	flag.UintVar(&lineLength, "l", 20, "length of each line, excluding spaces")
	flag.StringVar(&rowCode, "r", "hH", "code for rows to include, consult readme")
	flag.StringVar(&xtraChars, "x", "", "extra characters to practice")
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

	keys = make([]byte, 0)
	keys = append(keys, xtraChars...)
	if strings.Contains(rowCode, "n") {
		keys = append(keys, numLower...)
	}
	if strings.Contains(rowCode, "N") {
		keys = append(keys, numUpper...)
	}
	if strings.Contains(rowCode, "t") {
		keys = append(keys, topLower...)
	}
	if strings.Contains(rowCode, "T") {
		keys = append(keys, topUpper...)
	}
	if strings.Contains(rowCode, "h") {
		keys = append(keys, homeLower...)
	}
	if strings.Contains(rowCode, "H") {
		keys = append(keys, homeUpper...)
	}
	if strings.Contains(rowCode, "b") {
		keys = append(keys, botLower...)
	}
	if strings.Contains(rowCode, "B") {
		keys = append(keys, botUpper...)
	}
	amtKeys := len(keys)
	if amtKeys == 0 {
		panic("You haven't specified any keys to practice, consult readme on -r flag")
	}

	for k = amtLines; k > 0; k-- {
		// Construct line.
		for i = 0; i < lineLength; i++ {
			line[i] = keys[rand.Intn(amtKeys)]
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
