package repository

import (
	"bookapi/feature/book/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

// GetAll implements domain.Repository
func (rq *repoQuery) GetAll() ([]domain.Basic, error) {
	var resQry []Book
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil

}

// GetBook implements domain.Repository
func (rq *repoQuery) GetBook(ID uint) (domain.Basic, error) {
	var resQry Book
	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
		return domain.Basic{}, err
	}
	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}

// Insert implements domain.Repository
func (rq *repoQuery) Insert(newBook domain.Basic) (domain.Basic, error) {
	var cnv Book
	cnv = FromDomain(newBook)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Basic{}, err
	}
	// selesai dari DB
	newBook = ToDomain(cnv)
	return newBook, nil
}

// Update implements domain.Repository
func (rq *repoQuery) Update(updatedBook domain.Basic, ID uint) (domain.Basic, error) {
	var new Book
	if err := rq.db.First(&new, "ID = ?", ID).Error; err != nil {
		return domain.Basic{}, err
	}

	new.Judul = updatedBook.Judul
	new.Author = updatedBook.Author

	if err := rq.db.Save(&new).Error; err != nil {
		return domain.Basic{}, err
	}
	// selesai dari DB
	reUp := ToDomain(new)
	return reUp, nil
}

// Delete implements domain.Repository
func (rq *repoQuery) Delete(ID uint) (domain.Basic, error) {
	var resQry Book
	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
		log.Error(err.Error())
		return ToDomain(resQry), err
	}

	if err := rq.db.Delete(&resQry).Error; err != nil {
		log.Error(err.Error())
		return ToDomain(resQry), err
	}
	return ToDomain(resQry), nil

}
