// package server_test

// import (
// 	"context"
// 	"learngo-pockets/habits/api"
// 	"learngo-pockets/habits/internal/repository"
// 	"learngo-pockets/habits/internal/server"
// 	"net"
// 	"sync"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func TestIntegration(t *testing.T) {
// 	// Skip this test when running lightweight suites
// 	if testing.Short() {
// 		t.Skip()
// 	}
// 	grpcServ := newServer(t)
// 	listener, err := net.Listen("tcp", "")

// 	require.NoError(t, err)

// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		err = grpcServ.Serve(listener)
// 		require.NoError(t, err)
// 	}()
// 	defer func() {
// 		// terminate the GRPC server
// 		grpcServ.Stop()
// 		// if no error was caught, end the test
// 		wg.Wait()
// 	}()

// 	// create client
// 	habitsCli, err := newClient(t, listener.Addr().String())
// 	require.NoError(t,err)
// }

// func newClient(t *testing.T, serverAddress string) (api.HabitsClient, error) {
// 	t.Helper()
// 	creds := grpc.WithTransportCredentials(
// 		insecure.NewCredentials(),
// 	)

// 	conn, err := grpc.Dial(serverAddress, creds)
// 	if err !=nil {
// 		return nil, err
// 	}

// 	return api.NewHabitsClient(conn), nil
// }

// func newServer(t *testing.T) *grpc.Server {
// 	t.Helper()
// 	s := server.New(repository.New(t), t)

// 	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(timerInterceptor(s.lgr)))
// 	api.RegisterHabitsServer(grpcServer, s)
// 	reflection.Register(grpcServer)

// 	return grpcServer
// }

// func listHabitsMatches(t *testing.T, habitsCli api.HabitsClient, expected []*api.Habit) {
// 	list, err := habitsCli.ListHabits(context.Background(), &api.ListHabitsRequest{})
// 	require.NoError(t, err)

// 	for i := range list.Habits {
// 		assert.NotEqual(t, "", list.Habits[i].Id)
// 		list.Habits[i].Id = ""
// 	}
// 	assert.Equal(t, list.Habits, expected)
// }