package delivery

import (
	"bookapi/feature/book/domain"
	"strconv"

	//"bookapi/feature/book/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := bookHandler{srv: srv}
	e.GET("/books", handler.ShowAllBook())
	e.POST("/books", handler.AddBook())
	e.GET("/books/:id", handler.ThisBook())
	e.PUT("/books/:id", handler.EditBook())
	e.DELETE("/books/:id", handler.DeleteBook())

}

func (bs *bookHandler) AddBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddBookFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := bs.srv.AddBook(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}

}

func (bs *bookHandler) ShowAllBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bs.srv.AllBook()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all user", ToResponse(res, "all")))
	}
}

func (bs *bookHandler) ThisBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, err := strconv.Atoi(c.Param("id"))
		res, err := bs.srv.ThisBook(uint(ID))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Berhasil menemukan buku", res))
	}
}

func (bs *bookHandler) EditBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		//var update domain.Basic
		ID, err := strconv.Atoi(c.Param("id"))

		var newEdit EditBookFormat
		if err := c.Bind(&newEdit); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		update := domain.Basic{Judul: newEdit.Judul, Author: newEdit.Author}
		res, err := bs.srv.EditBook(update, uint(ID))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Berhasil edit buku", ToResponse(res, "edit")))
	}
}

func (bs *bookHandler) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, err := strconv.Atoi(c.Param("id"))
		if _, err = bs.srv.DeleteBook(uint(ID)); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Berhasil hapus buku", " "))
	}
}
