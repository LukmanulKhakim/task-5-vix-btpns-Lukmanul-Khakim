package delivery

import "bookapi/feature/book/domain"

type AddBookFormat struct {
	Judul  string `json:"judul" form:"judul"`
	Author string `json:"author" form:"author"`
}

type EditBookFormat struct {
	Judul  string `json:"judul" form:"judul"`
	Author string `json:"author" form:"author"`
}

func ToDomain(i interface{}) domain.Basic {
	switch i.(type) {
	case AddBookFormat:
		cnv := i.(AddBookFormat)
		return domain.Basic{Judul: cnv.Judul, Author: cnv.Author}
	case EditBookFormat:
		cnv := i.(EditBookFormat)
		return domain.Basic{Judul: cnv.Judul, Author: cnv.Author}
	}
	return domain.Basic{}
}
