package gordle

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"slices"
	"strings"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

// New returns a Game variable, which can be used to Play!
func New(reader io.Reader, wordList []string, maxAttempts int) (*Game, error) {
	if len(wordList) == 0 {
		return nil, CorpusError("corpus is empty")
	}

	g := &Game{
		reader:      bufio.NewReader(reader),
		solution:    []rune(strings.ToUpper(wordList[rand.Intn(len(wordList))])),
		maxAttempts: maxAttempts,
	}

	return g, nil
}

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("enter a %d-character guess:\n", len(g.solution))

	for {
		playerInput, _, err := g.reader.ReadLine()

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr,
				"Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s.\n", err.Error())
		} else {
			return guess
		}
	}
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	// ask for a valid word
	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()

		fb := computeFeedback(guess, g.solution)

		fmt.Println(fb.String())

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉You won! You found it in %d guess(es)! The word was: %s.\n", currentAttempt, string(g.solution))
			return
		}
	}

	fmt.Printf("😞You've lost! The solution was: %s. \n", string(g.solution))
}

// returned when the guess has the wrong number of characters
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

// validateGuess ensures the guess is valid enough
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}
	return nil
}

// splitToUppercaseCharacters is a naive implementation to turn a string into a list of characters
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}
