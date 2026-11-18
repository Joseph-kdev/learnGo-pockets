package habit_test

import (
	"context"
	"fmt"
	"learngo-pockets/habits/internal/habit"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mocklist is a mock for FindAll method response
type Mocklist struct {
	Items []habit.Habit
	Err   error
}

// FindAll is a mock which returns the passed list of items and error
func (list Mocklist) FindAll(context.Context) ([]habit.Habit, error) {
	return list.Items, list.Err
}

func TestListHabits(t *testing.T) {
	dbErr := fmt.Errorf("db unavailable")
	habits := []habit.Habit{{}, {}}

	tests := map[string]struct {
		db             Mocklist
		expectedErr    error
		expectedHabits []habit.Habit
	}{
		"empty": {
			db: Mocklist{
				Items: nil, Err: nil},
			expectedErr:    nil,
			expectedHabits: nil,
		},
		"2 items": {
			db:             Mocklist{Items: habits, Err: nil},
			expectedErr:    nil,
			expectedHabits: habits,
		},
		"error case": {
			db:             Mocklist{Items: nil, Err: dbErr},
			expectedErr:    dbErr,
			expectedHabits: nil,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := habit.ListAllHabits(
				context.Background(), tc.db)
			assert.ErrorIs(t, err, tc.expectedErr)
			assert.ElementsMatch(t, tc.expectedHabits, got)
		})
	}
}
