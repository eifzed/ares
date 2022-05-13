package commonerr

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewError(errorCode codes.Code, errorDesc string) error {
	return status.Errorf(errorCode, errorDesc)
}

func ErrorBadRequest(errorDesc string) error {
	return NewError(codes.InvalidArgument, errorDesc)
}

func ErrorAlreadyExist(errorDesc string) error {
	return NewError(codes.Aborted, errorDesc)
}

func ErrorUnauthorized(errorDesc string) error {
	return NewError(codes.Unauthenticated, errorDesc)
}

func ErrorForbidden(errorDesc string) error {
	return NewError(codes.PermissionDenied, errorDesc)
}

func ErrorNotFound(errorDesc string) error {
	return NewError(codes.NotFound, errorDesc)
}

func SetError(err error) error {
	stat, ok := status.FromError(err)
	if ok {
		return NewError(stat.Code(), stat.Message())
	}
	return NewError(codes.Internal, err.Error())
}
