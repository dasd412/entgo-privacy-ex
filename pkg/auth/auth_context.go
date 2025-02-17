package auth

import (
	"context"
	"errors"
	"strconv"
)

type userIdContextKey string

const userIDKey userIdContextKey = "userID"

func WithUserId(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, userIDKey, userId)
}

func GetUserId(ctx context.Context) (int, error) {
	userIdString, ok := ctx.Value(userIDKey).(string)

	if !ok {
		return -1, errors.New("user ID not found in context")
	}

	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return -1, err
	}

	return userId, nil
}
