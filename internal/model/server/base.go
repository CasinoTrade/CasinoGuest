package server

import "context"

type Base interface {
	Start()
	Stop()
	Ping(context.Context) int
}
