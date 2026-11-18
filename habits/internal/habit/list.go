package habit

import "context"

type habitList interface {
	FindAll(ctx context.Context) ([]Habit, error)
}

func ListAllHabits(ctx context.Context, db habitList) ([]Habit, error) {
	habits, err := db.FindAll(ctx)
	if err != nil {
		return []Habit{}, err
	}

	return habits, nil
}