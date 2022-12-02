package delivery

import "bookapi/feature/book/domain"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type AddBookResponse struct {
	ID     uint   `json:"id"`
	Judul  string `json:"judul"`
	Author string `json:"author"`
}

type EditBookResponse struct {
	Judul  string `json:"judul"`
	Author string `json:"author"`
}

func ToResponse(basic interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := basic.(domain.Basic)
		res = AddBookResponse{ID: cnv.ID, Judul: cnv.Judul, Author: cnv.Author}
	case "all":
		var arr []AddBookResponse
		cnv := basic.([]domain.Basic)
		for _, val := range cnv {
			arr = append(arr, AddBookResponse{ID: val.ID, Judul: val.Judul, Author: val.Author})
		}
		res = arr
	case "edit":
		cnv := basic.(domain.Basic)
		res = EditBookResponse{Judul: cnv.Judul, Author: cnv.Author}
	}

	return res
}
