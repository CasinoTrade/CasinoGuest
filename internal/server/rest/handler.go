package rest

import (
	"github.com/CasinoTrade/CasinoGuest/internal/server"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	base *server.Casino
}

func newHandler(base *server.Casino) *handler {
	return &handler{
		base: base,
	}
}

func (h *handler) Ping(c *fiber.Ctx) error {
	res := h.base.Ping(c.Context())
	return c.JSON(fiber.Map{
		"Number": res,
	})
}
