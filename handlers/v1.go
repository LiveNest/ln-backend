package feature

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

// @host localhost:8080
// @BasePath /dbms/v1
func RegisterHandlerV1(db *gorm.DB) *fiber.App {
	router := fiber.New()
	_ = router.Group("/dbms/v1")
	return router
}
