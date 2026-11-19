package repository

import (
	"context"
	"fmt"
	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/isoweek"
	"sort"
	"sync"
	"time"
)

// Logger used by the repository
type Logger interface {
	Logf(format string, args ...any)
}

// ticksPerWeek holds all the timestamps for a given week number
type ticksPerWeek map[isoweek.ISO8601][]time.Time

// HabitRepository holds all the current habits
type HabitRepository struct {
	mutex   sync.Mutex
	storage map[habit.ID]habit.Habit
	lgr     Logger
	habits  map[habit.ID]habit.Habit
	ticks   map[habit.ID]ticksPerWeek
}

// New creates an empty habit repository
func New(lgr Logger) *HabitRepository {
	return &HabitRepository{
		lgr:     lgr,
		storage: make(map[habit.ID]habit.Habit),
	}
}

// Add inserts for the first time a habit in memory
func (hr *HabitRepository) Add(_ context.Context, habit habit.Habit) error {
	hr.lgr.Logf("Adding a habit...")

	// Lock the writing of the habit
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	hr.storage[habit.ID] = habit

	return nil
}

// findAll returns all habits
func (hr *HabitRepository) FindAll(_ context.Context) ([]habit.Habit, error) {
	hr.lgr.Logf("Listing habits...")

	// Lock the reading and the writing of the habits
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	habits := make([]habit.Habit, 0)
	for _, h := range hr.storage {
		habits = append(habits, h)
	}

	// Ensure the output is deterministic by sorting the habits
	sort.Slice(habits, func(i, j int) bool {
		return habits[i].CreationTime.Before(habits[j].CreationTime)
	})

	return habits, nil
}

func (hr *HabitRepository) Find(_ context.Context, id habit.ID) (habit.Habit, error) {
	hr.lgr.Logf("Finding a habit...")
	h, found := hr.habits[id]
	if !found {
		return habit.Habit{}, fmt.Errorf("habit %q not registered: %w", id, ErrNotFound)
	}

	return h, nil
}

// AddTick inserts a new event for a habit in memory
func (hr *HabitRepository) AddTick(_ context.Context, id habit.ID, t time.Time) error {
	hr.lgr.Logf("Adding a tick...")

	// Lock the reading and the writing of the habits
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	_, found := hr.ticks[id]
	if !found {
		hr.ticks[id] = make(ticksPerWeek)
	}

	w := isoweek.At(t)

	ticks, found := hr.ticks[id][w]
	if !found {
		ticks = make([]time.Time, 0, 1)
	}
	hr.ticks[id][w] = append(ticks, t)

	return nil
}