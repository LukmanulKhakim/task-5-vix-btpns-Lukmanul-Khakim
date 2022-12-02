package repository

import (
	"bookapi/feature/book/domain"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Judul  string
	Author string
	IdUser uint
}

func FromDomain(dk domain.Basic) Book {
	return Book{
		Model:  gorm.Model{ID: dk.ID},
		Judul:  dk.Judul,
		Author: dk.Author,
		IdUser: dk.IdUser,
	}
}

func ToDomain(b Book) domain.Basic {
	return domain.Basic{
		ID:     b.ID,
		Judul:  b.Judul,
		Author: b.Author,
		IdUser: b.IdUser,
	}
}

func ToDomainArray(au []Book) []domain.Basic {
	var res []domain.Basic
	for _, val := range au {
		res = append(res, domain.Basic{ID: val.ID, Judul: val.Judul, Author: val.Author, IdUser: val.IdUser})
	}

	return res
}
