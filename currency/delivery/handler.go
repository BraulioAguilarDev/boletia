package handler

import (
	"boletia/currency"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Usecase currency.Usecase
}

func NewHandler(usecase currency.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}

func (h *Handler) Get(c *fiber.Ctx) error {

	// h.UseCase.Get()

	return nil
}
