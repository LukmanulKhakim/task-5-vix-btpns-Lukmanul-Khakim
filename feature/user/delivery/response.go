package delivery

import "bookapi/feature/user/domain"

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

type RegisterResponse struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
	HP   string `json:"hp"`
}

type LoginResponse struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = RegisterResponse{ID: cnv.ID, Nama: cnv.Nama, HP: cnv.HP}
	case "login":
		cnv := core.(domain.Core)
		res = LoginResponse{ID: cnv.ID, Nama: cnv.Nama}
	case "all":
		var arr []RegisterResponse
		cnv := core.([]domain.Core)
		for _, val := range cnv {
			arr = append(arr, RegisterResponse{ID: val.ID, Nama: val.Nama, HP: val.HP})
		}
		res = arr
	}

	return res
}
