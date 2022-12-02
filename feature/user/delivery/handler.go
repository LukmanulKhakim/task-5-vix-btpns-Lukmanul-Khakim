package delivery

import (
	"bookapi/feature/user/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.POST("/register", handler.Register())
	e.POST("/login", handler.Login())
	e.PUT("/users/:id", handler.UpdateProfile())
	e.GET("/users/:id", handler.Profile())
	e.GET("/users", handler.ShowAllUser())
	e.DELETE("/users", handler.DeleteUser())
}

// registrasi add user
func (us *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.AddUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}

}

func (us *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		//cnv := ToDomain(input)
		res, err := us.srv.Login(domain.Core{Nama: input.Nama, Password: input.Password})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("berhasil login", ToResponse(res, "login")))
	}
}

// update user
func (us *userHandler) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var up UpdateFormat
		ID, err := strconv.Atoi(c.Param("id"))
		if err := c.Bind(&up); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Cant bind update data"))
		}

		cnv := ToDomain(up)
		res, err := us.srv.UpdateProfile(cnv, uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))

		}
		return c.JSON(http.StatusAccepted, SuccessResponse("berhasil update", res))
	}
}

// ambil ID User
func (us *userHandler) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID, err := strconv.Atoi(c.Param("id"))
		res, err := us.srv.Profile(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("berhasil menemukan data", res))
	}
}

// ambil semua data
func (us *userHandler) ShowAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := us.srv.ShowAllUser()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all user", ToResponse(res, "all")))
	}
}

// hapus data
func (us *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var del DeleteFormat
		if err := c.Bind(&del); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Cant bind update data"))
		}

		cnv := ToDomain(del)
		res, err := us.srv.DeleteUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusAccepted, SuccessResponse("berhasil hapus data", res))
	}
}
