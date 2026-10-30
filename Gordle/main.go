package main

import (
	"bufio"
	"fmt"
	"learngo-pockets/gordle/corpus"
	"learngo-pockets/gordle/gordle"
	"os"
)

const maxAttempts = 6

func main() {
	corpus, err := corpus.ReadCorpus("corpus/english.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
		return
	}

	//create the game
	g, err := gordle.New(bufio.NewReader(os.Stdin), corpus, maxAttempts)
	if err!=nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start game: %s", err)
		return
	}

	// Run the game
	g.Play()
}