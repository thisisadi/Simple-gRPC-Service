package main

import (
	"context"
	"log"
	"net"

	"github.com/thisisadi/Total-Corp-Task/records"
	"github.com/thisisadi/Total-Corp-Task/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	user.UnimplementedUserServiceServer
}

func (s *UserServiceServer) GetUserById(ctx context.Context, request *user.UserRequest) (*user.User, error) {

	userID := request.Id
	for _, u := range records.Users {
		if u.Id == userID {
			return u, nil
		}
	}

	// If no user is found with the given ID, return an error
	return nil, status.Errorf(codes.NotFound, "User with ID %d not found", userID)
}

func (s *UserServiceServer) GetUsersByIds(ctx context.Context, request *user.UserIdsRequest) (*user.UsersResponse, error) {
	userIDs := request.Ids
	var matchedUsers []*user.User

	// Iterate over the IDs and find matching users
	for _, userID := range userIDs {
		for _, u := range records.Users {
			if u.Id == userID {
				matchedUsers = append(matchedUsers, u)
				break
			}
		}
	}

	// Create a UsersResponse with the matched users
	response := &user.UsersResponse{
		Users: matchedUsers,
	}

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("unable to listen : %s", err)
	}
	println("server running on port 3000")
	serverRegistrar := grpc.NewServer()
	service := &UserServiceServer{}

	user.RegisterUserServiceServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("server stopped : %s", err)
	}
}
