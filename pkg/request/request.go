// Package request implements request handling utilities.
package request

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Handle simplifies the request-response cycle by executing concrete functions and passed arguments.
func Handle(c *fiber.Ctx, req interface{}, timeout time.Duration, respondWithErr func(error) Response, handler func(ctx context.Context, req any) (any, error)) error {
	if err := binder(c, req); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(c.UserContext(), timeout)
	defer cancel()

	data, err := handler(ctx, req)
	if err != nil {
		response := respondWithErr(err)
		return c.Status(response.Status).JSON(response)
	}

	status := c.Response().StatusCode()

	response := NewResponse(status, "Success", data, nil)
	return c.Status(status).JSON(response)
}

// binder binds the data from the request to a pointer struct.
func binder(c *fiber.Ctx, req any) error {
	if req == nil {
		return nil
	}

	if err := c.QueryParser(req); err != nil {
		resp := bindQueryErr(err)
		return c.Status(resp.Status).JSON(resp)
	}

	if err := c.ParamsParser(req); err != nil {
		resp := bindPathParamErr(err)
		return c.Status(resp.Status).JSON(resp)
	}

	if string(c.Request().Header.ContentType()) == "application/json" {
		if err := c.BodyParser(req); err != nil {
			resp := bindJSONErr(err)
			return c.Status(resp.Status).JSON(resp)
		}
	}

	return nil
}
