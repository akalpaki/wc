package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// parse flags
	var modeBytes bool
	var modeLines bool
	var modeWords bool
	var modeChars bool
	flag.BoolVar(&modeBytes, "c", false, "count mode: prints the number of bytes in input")
	flag.BoolVar(&modeLines, "l", false, "line mode: prints the number of lines in input")
	flag.BoolVar(&modeWords, "w", false, "word mode: prints the number of words in input")
	flag.BoolVar(&modeChars, "m", false, "character mode: prints the number of characters in input")
	flag.Parse()

	flags := []bool{modeBytes, modeLines, modeWords, modeChars}
	if !checkNoFlags(flags) {
		// if no flags, use default setup -c -l -w
		modeBytes = true
		modeLines = true
		modeWords = true
	}

	// get reader (TODO: implement stdin)
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("STDIN not implemented.")
		os.Exit(1)
	}
	rdr, err := os.Open(filename)
	if err != nil {
		fmt.Println("unable to open file")
		os.Exit(1)
	}
	defer rdr.Close()

	// read data from reader
	// create buffers
	var numBytes int
	var numLines int
	var numWords int
	var numChars int
	var output string

	if modeBytes {
		bufByte := bufio.NewScanner(rdr)
		bufByte.Split(bufio.ScanBytes)
		for bufByte.Scan() {
			numBytes += len(bufByte.Bytes())
		}
		output += fmt.Sprintf(" %d ", numBytes)
	}

	if modeLines {
		resetFile(rdr)
		bufLine := bufio.NewScanner(rdr)
		bufLine.Split(bufio.ScanLines)
		for bufLine.Scan() {
			numLines += 1
		}
		output += fmt.Sprintf(" %d ", numLines)
	}

	if modeWords {
		resetFile(rdr)
		bufWord := bufio.NewScanner(rdr)
		bufWord.Split(bufio.ScanWords)
		for bufWord.Scan() {
			numWords += 1
		}
		output += fmt.Sprintf(" %d ", numWords)
	}

	if modeChars {
		resetFile(rdr)
		bufChar := bufio.NewScanner(rdr)
		bufChar.Split(bufio.ScanRunes)
		for bufChar.Scan() {
			numChars += 1
		}
		output += fmt.Sprintf(" %d ", numChars)
	}

	fmt.Printf("%s %s\n", output, filename)
}

func resetFile(f *os.File) {
	_, err := f.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Printf("failed to reset file: %s", err.Error())
		os.Exit(1)
	}
}

func checkNoFlags(flags []bool) bool {
	for _, flag := range flags {
		if flag {
			return true
		}
	}
	return false
}
