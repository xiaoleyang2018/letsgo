package logices

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/qinhao/letsgo/models"
	"net/http"
)
type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
type CustomValidator struct {
validator *validator.Validate
}
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func Users(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest,"request err")
		return
	}
	if err = c.Validate(u); err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}
	//落入数据库

	uid, err := models.SaveUser(&models.User{
		Name: u.Name,
		Email: u.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}


	return c.JSON(http.StatusOK, uid)
}