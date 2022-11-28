package asanihelper_test

import (
	"net/http/httptest"
	"testing"

	asanihelper "github.com/asani-co-id/go-asani-helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFiberReponsesSuccess(t *testing.T) {
	//init response success
	assert := assert.New(t)

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		// Return simple string as response
		return asanihelper.FiberReponses(c, fiber.StatusOK, "success", "success test", nil)
	})

	newResponses := asanihelper.Response{
		Status:  "success",
		Message: "success test",
		Data:    nil,
	}

	req := httptest.NewRequest("GET", "/hello", nil)
	resp, _ := app.Test(req, 1)

	assert.Equalf(fiber.StatusOK, fiber.StatusOK, resp.Status, newResponses.Message)
}

func TestFiberReponsesBadRequest(t *testing.T) {
	//init response success
	assert := assert.New(t)

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		// Return simple string as response
		return asanihelper.FiberReponses(c, fiber.StatusBadRequest, "bad request", "bad request test", nil)
	})

	newResponses := asanihelper.Response{
		Status:  "bad request",
		Message: "bad request test",
		Data:    nil,
	}

	req := httptest.NewRequest("GET", "/hello", nil)
	resp, _ := app.Test(req, 1)

	assert.Equalf(fiber.StatusBadRequest, fiber.StatusBadRequest, resp.Status, newResponses.Message)
}

func TestFiberReponsesNotFound(t *testing.T) {
	//init response success
	assert := assert.New(t)

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		// Return simple string as response
		return asanihelper.FiberReponses(c, fiber.StatusNotFound, "not found", "not found test", nil)
	})

	newResponses := asanihelper.Response{
		Status:  "not found",
		Message: "not found test",
		Data:    nil,
	}

	req := httptest.NewRequest("GET", "/hello", nil)
	resp, _ := app.Test(req, 1)

	assert.Equalf(fiber.StatusNotFound, fiber.StatusNotFound, resp.Status, newResponses.Message)
}

func TestFiberReponsesInternalServerError(t *testing.T) {
	//init response success
	assert := assert.New(t)

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		// Return simple string as response
		return asanihelper.FiberReponses(c, fiber.StatusInternalServerError, "internal server error", "internal server error test", nil)
	})

	newResponses := asanihelper.Response{
		Status:  "internal server error",
		Message: "internal server error test",
		Data:    nil,
	}

	req := httptest.NewRequest("GET", "/hello", nil)
	resp, _ := app.Test(req, 1)

	assert.Equalf(fiber.StatusInternalServerError, fiber.StatusInternalServerError, resp.Status, newResponses.Message)
}

func TestFiberReponsesConflict(t *testing.T) {
	//init response success
	assert := assert.New(t)

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		// Return simple string as response
		return asanihelper.FiberReponses(c, fiber.StatusConflict, "Conflict", "Conflict test", nil)
	})

	newResponses := asanihelper.Response{
		Status:  "Conflict",
		Message: "Conflict test",
		Data:    nil,
	}

	req := httptest.NewRequest("GET", "/hello", nil)
	resp, _ := app.Test(req, 1)

	assert.Equalf(fiber.StatusConflict, fiber.StatusConflict, resp.Status, newResponses.Message)
}

func TestFiberReponsesUnauthorized(t *testing.T) {
	//init response success
	assert := assert.New(t)

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		// Return simple string as response
		return asanihelper.FiberReponses(c, fiber.StatusUnauthorized, "Unauthorized", "Unauthorized test", nil)
	})

	newResponses := asanihelper.Response{
		Status:  "Unauthorized",
		Message: "Unauthorized test",
		Data:    nil,
	}

	req := httptest.NewRequest("GET", "/hello", nil)
	resp, _ := app.Test(req, 1)

	assert.Equalf(fiber.StatusUnauthorized, fiber.StatusUnauthorized, resp.Status, newResponses.Message)
}

func TestFiberReponsesForbidden(t *testing.T) {
	//init response success
	assert := assert.New(t)

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		// Return simple string as response
		return asanihelper.FiberReponses(c, fiber.StatusForbidden, "Forbidden", "Forbidden test", nil)
	})

	newResponses := asanihelper.Response{
		Status:  "Forbidden",
		Message: "Forbidden test",
		Data:    nil,
	}

	req := httptest.NewRequest("GET", "/hello", nil)
	resp, _ := app.Test(req, 1)

	assert.Equalf(fiber.StatusForbidden, fiber.StatusForbidden, resp.Status, newResponses.Message)
}

func TestFiberReponsesTooManyRequest(t *testing.T) {
	//init response success
	assert := assert.New(t)

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		// Return simple string as response
		return asanihelper.FiberReponses(c, fiber.StatusTooManyRequests, "TooManyRequests", "TooManyRequests test", nil)
	})

	newResponses := asanihelper.Response{
		Status:  "TooManyRequests",
		Message: "TooManyRequests test",
		Data:    nil,
	}

	req := httptest.NewRequest("GET", "/hello", nil)
	resp, _ := app.Test(req, 1)

	assert.Equalf(fiber.StatusTooManyRequests, fiber.StatusTooManyRequests, resp.Status, newResponses.Message)
}

func TestFiberReponsesEntitiesTooLarget(t *testing.T) {
	//init response success
	assert := assert.New(t)

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		// Return simple string as response
		return asanihelper.FiberReponses(c, fiber.StatusRequestEntityTooLarge, "EntityTooLarge", "EntityTooLarge test", nil)
	})

	newResponses := asanihelper.Response{
		Status:  "EntityTooLarge",
		Message: "EntityTooLarge test",
		Data:    nil,
	}

	req := httptest.NewRequest("GET", "/hello", nil)
	resp, _ := app.Test(req, 1)

	assert.Equalf(fiber.StatusRequestEntityTooLarge, fiber.StatusRequestEntityTooLarge, resp.Status, newResponses.Message)
}

func TestFiberReponsesNoContent(t *testing.T) {
	//init response success
	assert := assert.New(t)

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		// Return simple string as response
		return asanihelper.FiberReponses(c, fiber.StatusNoContent, "StatusNoContent", "StatusNoContent test", nil)
	})

	newResponses := asanihelper.Response{
		Status:  "StatusNoContent",
		Message: "StatusNoContent test",
		Data:    nil,
	}

	req := httptest.NewRequest("GET", "/hello", nil)
	resp, _ := app.Test(req, 1)

	assert.Equalf(fiber.StatusNoContent, fiber.StatusNoContent, resp.Status, newResponses.Message)
}
