package habit

import (
	"context"
	"strings"
	"time"
	"uuid"
)

// validateAndCompleteHabit fills the habit with values that we want in
// our database
// Returns InvalidInputError
func validateAndCompleteHabit(h Habit) (Habit, error) {
	// name cannot be empty
	h.Name = Name(strings.TrimSpace(string(h.Name)))
	if h.Name == "" {
		return Habit{}, InvalidInputError{
			field: "name", reason: "cannot be empty",
		}
	}

	if h.WeeklyFrequency == 0 {
		h.WeeklyFrequency = 1
	}

	if h.ID == "" {
		h.ID = ID(uuid.New().String())
	}

	if h.CreationTime.Equal(time.Time{}) {
		h.CreationTime = time.Now()
	}

	return h, nil
}

// create validates the Habit, saves it and returns it
func Create(_ context.Context, h Habit) (Habit, error) {
	h, err := validateAndCompleteHabit(h)
	if err != nil {
		return Habit{}, err
	}
	// add the habit to data storage
	return h, nil
}