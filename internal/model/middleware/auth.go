package middleware

import "context"

type MiddlewareInterface interface {
	GRPCAuthMiddlewareInterface
}

type GRPCAuthMiddlewareInterface interface {
	AuthHandler(ctx context.Context) (context.Context, error)
}
