package main

import (
	"github.com/eifzed/ares/internal/handler"
	"github.com/eifzed/ares/internal/model/middleware"
)

type modules struct {
	GRPCHandler    handler.GRPCHandler
	GRPCMiddleware middleware.MiddlewareInterface
}

func getNewModules(mod *modules) *modules {
	return mod
}
