package services

import (
	"bookapi/feature/book/domain"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
)

type bookService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &bookService{
		qry: repo,
	}
}

// AddBook implements domain.Service
func (bs *bookService) AddBook(newBook domain.Basic) (domain.Basic, error) {
	res, err := bs.qry.Insert(newBook)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Basic{}, errors.New("rejected from database")
		}

		return domain.Basic{}, errors.New("some problem on database")
	}

	return res, nil
}

// AllBook implements domain.Service
func (bs *bookService) AllBook() ([]domain.Basic, error) {
	res, err := bs.qry.GetAll()
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("no data")
		}
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}

	return res, nil

}

// ThisBook implements domain.Service
func (bs *bookService) ThisBook(ID uint) (domain.Basic, error) {
	res, err := bs.qry.GetBook(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Basic{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Basic{}, errors.New("no data")
		}
	}

	return res, nil
}

// EditBook implements domain.Service
func (bs *bookService) EditBook(updatedBook domain.Basic, ID uint) (domain.Basic, error) {
	res, err := bs.qry.Update(updatedBook, ID)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Basic{}, errors.New("rejected from database")
		}
	}
	return res, nil
}

// DeleteBook implements domain.Service
func (bs *bookService) DeleteBook(ID uint) (domain.Basic, error) {
	res, err := bs.qry.Delete(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "column") {
			return domain.Basic{}, errors.New("rejected from database")
		}
	}
	return res, nil
}
