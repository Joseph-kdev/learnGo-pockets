package server

import (
	"context"
	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListHabits is the endpoint that lists all habits added
func (s *Server) ListHabits(ctx context.Context, request *api.ListHabitsRequest) (*api.ListHabitsResponse, error) {
	habits, err := habit.ListAllHabits(ctx, s.db)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot list habits: %s", err.Error())
	}

	s.lgr.Logf("Fetched %d habits", len(habits))

	return &api.ListHabitsResponse{
		Habits: toAPIHabits(habits),
	}, nil
}

func toAPIHabits(habits []habit.Habit) []*api.Habit {
	apiHabits := make([]*api.Habit, 0, len(habits))
	for _, h := range habits {
		apiHabits = append(apiHabits, &api.Habit{
			Name:            string(h.Name),
			WeeklyFrequency: int32(h.WeeklyFrequency),
		})
	}

	return apiHabits
}
