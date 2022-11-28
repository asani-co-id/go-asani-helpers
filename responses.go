package asanihelper

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  interface{} `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func (m *Response) successContext(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&m)
}

func (m *Response) badRequestContext(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&m)
}

func (m *Response) notFoundContext(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(&m)
}

func (m *Response) internalServerErrorContext(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(&m)
}

func (m *Response) conflictErrorContext(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusConflict).JSON(&m)
}

func (m *Response) unauthorizedContext(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(&m)
}

func (m *Response) forbiddenContext(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusForbidden).JSON(&m)
}

func (m *Response) tooManyRequestContext(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusTooManyRequests).JSON(&m)
}

func (m *Response) entitiesTooLargeContext(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusRequestEntityTooLarge).JSON(&m)
}

func (m *Response) noContentContext(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNoContent).JSON(&m)
}

func FiberReponses(ctx *fiber.Ctx, code int, status interface{}, message interface{}, data interface{}) error {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	switch code {
	case http.StatusOK:
		return response.successContext(ctx)
	case http.StatusBadRequest:
		return response.badRequestContext(ctx)
	case http.StatusNotFound:
		return response.notFoundContext(ctx)
	case http.StatusInternalServerError:
		return response.internalServerErrorContext(ctx)
	case http.StatusConflict:
		return response.conflictErrorContext(ctx)
	case http.StatusUnauthorized:
		return response.unauthorizedContext(ctx)
	case http.StatusForbidden:
		return response.forbiddenContext(ctx)
	case http.StatusTooManyRequests:
		return response.tooManyRequestContext(ctx)
	case http.StatusRequestEntityTooLarge:
		return response.entitiesTooLargeContext(ctx)
	default:
		return response.noContentContext(ctx)
	}
}
