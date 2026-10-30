package gordle

import (
	"fmt"
	"os"
	"strings"
)

// hint describes the validity of a character in a word
type hint byte

// feedback is a list of hints one per character of the word
type feedback []hint

func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// String implements the Stringer interface
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬛"
	case wrongPosition:
		return "🟠"
	case correctPosition:
		return "💚"
	default:
		return "💔"
	}
}

// computeFeedback verifies every character of the guess against the solution
func computeFeedback(guess, solution []rune) feedback {
	// initialise holders for marks
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error! Guess and solution"+
			" have different lengths: %d vs %d", len(guess), len(solution))
		return result
	}

	// check for correct letters
	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}

	// look for letters in the wrong position
	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			// The character has already been marked, ignore it
			continue
		}

		for posInSolution, target := range solution {
			if used[posInSolution] {
				// The letter of the solution is already assigned
				// to a letter of the guess
				// skip to the next letter of the solution
				continue
			}
			if character == target {
				result[posInGuess] = wrongPosition
				used[posInSolution] = true
				// skip to the next letter of the guess
				break
			}
		}
	}
	return result
}

// Equal determines equality of two feedbacks.
func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}

	for index, value := range fb {
		if value != other[index] {
			return false
		}
	}
	return true
}
