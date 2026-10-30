package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input), string(tc.want), 0)

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word     []rune
		expected error
	}{
		"nominal": {
			word:     []rune("GUESS"),
			expected: nil,
		},
		"too long": {
			word:     []rune("POCKET"),
			expected: errInvalidWordLength,
		},
		"blank": {
			word:     []rune(""),
			expected: errInvalidWordLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(nil, string("GUESS"), 0)

			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, err)
			}
		})
	}
}

func TextComputeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {
			guess: "land",
			solution: "land",
			expectedFeedback: feedback{
				correctPosition,
				correctPosition,
				correctPosition,
				correctPosition,
			},
		},
		"double character": {
			guess: "small",
			solution: "small",
			expectedFeedback: feedback{
				correctPosition,
				correctPosition,
				correctPosition,
				correctPosition,
			},
		},
		"double character with wrong answer": {
			guess: "small",
			solution: "hello",
			expectedFeedback: feedback{
				absentCharacter,
				absentCharacter,
				absentCharacter,
				correctPosition,
				wrongPosition,
			},
		},
		"twp identical, but not in the right position": {
			guess: "hlleo",
			solution: "hello",
			expectedFeedback: feedback{
				correctPosition,
				wrongPosition,
				correctPosition,
				wrongPosition,
				correctPosition,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf(
					"guess: %q, got the wrong feedback, wanted %v, got %v",
					tc.guess, tc.expectedFeedback, fb)
			}
		})
	}
}
