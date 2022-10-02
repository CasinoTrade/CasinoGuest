package rest

import (
	"context"
	"net/http"

	model "github.com/CasinoTrade/CasinoGuest/internal/model/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CasinoREST struct {
	ctx    context.Context
	cancel context.CancelFunc

	adress string
	e      *echo.Echo
}

func (s *CasinoREST) Start() {
	e := echo.New()
	s.e = e

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Number int }{Number: 42})
	})

	e.Logger.Fatal(e.Start(s.adress))
}

func (s *CasinoREST) Stop() {
	if s.cancel != nil {
		defer s.cancel()
	}
	// NB: Shutdown waits for active connections, without interrupting them.
	// May be we should consider using Stop.
	s.e.Shutdown(context.Background())
}

func New(adress string) model.Casino {
	ctx, cancel := context.WithCancel(context.Background())
	return &CasinoREST{
		adress: adress,
		ctx:    ctx,
		cancel: cancel,
	}
}
