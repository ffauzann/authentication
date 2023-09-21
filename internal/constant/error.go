package constant

import (
	"errors"

	"google.golang.org/grpc/codes"
)

// Known gRPC errors
var (
	// Generic errors
	ErrNotFound = errors.New("Not found")
	ErrNoArg    = errors.New("No argument given")
	ErrInternal = errors.New("Internal error")

	// Specific errors
	ErrInvalidUsernamePassword = errors.New("Invalid username/password")
	ErrPasswordIsTooWeak       = errors.New("Password is too weak")
	ErrMalformedEmail          = errors.New("Malformed email")
	ErrInvalidUserIdType       = errors.New("Invalid user ID type")
	ErrUserNotFound            = errors.New("User not found")
	ErrUserIsBlocked           = errors.New("User is blocked")
)

// All client-safe errors goes here
var (
	MapGRPCErrCodes = map[error]codes.Code{
		// For HTTP mapping: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
		ErrInvalidUsernamePassword: codes.InvalidArgument,
		ErrMalformedEmail:          codes.InvalidArgument,
		ErrInvalidUserIdType:       codes.InvalidArgument,
		ErrPasswordIsTooWeak:       codes.FailedPrecondition,
		ErrNoArg:                   codes.FailedPrecondition,
		ErrNotFound:                codes.NotFound,
		ErrUserNotFound:            codes.NotFound,
		ErrUserIsBlocked:           codes.Aborted,
		ErrInternal:                codes.Internal,
	}
)
