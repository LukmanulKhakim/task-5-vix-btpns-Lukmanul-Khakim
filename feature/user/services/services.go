package services

import (
	"bookapi/feature/user/domain"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{
		qry: repo,
	}
}

// registrasi Add User
func (us *userService) AddUser(newUser domain.Core) (domain.Core, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("cannot encript password")
	}

	newUser.Password = string(generate)
	res, err := us.qry.Insert(newUser)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}

		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil
}

// Login implements domain.Service
func (us *userService) Login(newUser domain.Core) (domain.Core, error) {
	res, err := us.qry.GetUser(newUser)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}
	//resToken :=
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(newUser.Password))
	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("incorrect password")
	}
	return res, nil
}

// Update Data User
func (us *userService) UpdateProfile(updatedData domain.Core, ID uint) (domain.Core, error) {
	if updatedData.Password != "" {
		generate, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error(err.Error())
			return domain.Core{}, errors.New("cannot encript password")
		}
		updatedData.Password = string(generate)
	}
	//updatedData.Nama =

	res, err := us.qry.Update(updatedData, ID)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil
}

// ambil ID user
func (us *userService) Profile(ID uint) (domain.Core, error) {
	res, err := us.qry.Get(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}

	return res, nil
}

// ambil semua data
func (us *userService) ShowAllUser() ([]domain.Core, error) {
	res, err := us.qry.GetAll()
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

// DeleteUser implements domain.Service
func (us *userService) DeleteUser(deleteUser domain.Core) (domain.Core, error) {
	res, err := us.qry.Delete(deleteUser)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}

	return res, nil
}
