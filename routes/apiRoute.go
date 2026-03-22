package routers

import (
	svcHealthcheck "template_school/pkg/services/healthCheck"

	"github.com/gofiber/fiber/v3"
)

func APIRoute(app *fiber.App) {
	publicV1 := app.Group("/api/public/v1")
	privateV1 := app.Group("/api/private/v1")

	// HealthCheck
	publicV1.Get("/", svcHealthcheck.HealthCheck)
	privateV1.Get("/", svcHealthcheck.HealthCheck)
}
