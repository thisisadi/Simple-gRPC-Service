package utils

import (
	"context"

	"github.com/thisisadi/Total-Corp-Task/records"
	"github.com/thisisadi/Total-Corp-Task/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	user.UnimplementedUserServiceServer
}

func (s *UserServiceServer) GetUserById(ctx context.Context, request *user.UserRequest) (*user.User, error) {

	userID := request.Id
	u, ok := records.Users[int(userID)]
	if ok {
		return u, nil
	}

	// If no user is found with the given ID, return an error
	return nil, status.Errorf(codes.NotFound, "User with ID %d not found", userID)
}

func (s *UserServiceServer) GetUsersByIds(ctx context.Context, request *user.UserIdsRequest) (*user.UsersResponse, error) {
	userIDs := request.Ids
	var matchedUsers []*user.User

	for _, userID := range userIDs {
		u, ok := records.Users[int(userID)]
		if ok {
			matchedUsers = append(matchedUsers, u)
		}
	}

	response := &user.UsersResponse{
		Users: matchedUsers,
	}

	return response, nil
}
