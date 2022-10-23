package rest

import (
	model "github.com/CasinoTrade/CasinoGuest/internal/model/server"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	base model.Base
}

func newHandler(base model.Base) *handler {
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
