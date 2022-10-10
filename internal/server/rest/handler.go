package rest

import (
	"context"
	"net/http"

	"github.com/CasinoTrade/CasinoGuest/internal/server"
	"github.com/labstack/echo/v4"
)

type handler struct {
	base *server.Casino
}

func newHandler(base *server.Casino) *handler {
	return &handler{
		base: base,
	}
}

func (h *handler) Ping(c echo.Context) error {
	res := h.base.Ping(context.TODO())
	return c.JSON(http.StatusOK, struct{ Number int }{res})
}
