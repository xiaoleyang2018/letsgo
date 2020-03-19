package router

import (
	"github.com/go-playground/validator"
	"github.com/qinhao/letsgo/logices"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/go-playground/validator"
)

// Init initialize the route
func Init(logPath string, mode string) (*echo.Echo, error) {
	out := os.Stdout
	if mode == "prod" {
		f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		out = f
	}
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri},status=${status}\n",
		Output: out,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/hello", func(c echo.Context) error {
		return c.String(200, "Hello World")
	})
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/users",logices.Users )

	return e, nil
}
type CustomValidator struct {
	validator *validator.Validate
}
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}