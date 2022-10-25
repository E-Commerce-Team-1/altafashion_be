package delivery

import "altafashion_be/feature/users/domain"

func SuccessResponses(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponses(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

func SuccessDeleteResponses(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

type registerRespons struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type loginResponses struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type EditUserResponseFormat struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Location string `json:"location"`
	Profile  string `json:"profile"`
}
type GetUserResponseFormat struct {
	ID       uint   `json:"id"`
	Fullname string `json:"Fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
	Profile  string `json:"profile"`
}

func ToResponse(core interface{}, code string, token string) interface{} {
	var res interface{}

	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = registerRespons{ID: cnv.ID, Fullname: cnv.Fullname, Email: cnv.Email}
	case "login":
		cnv := core.(domain.Core)
		res = loginResponses{Fullname: cnv.Fullname, Email: cnv.Email, Token: token}
	case "edit":
		cnv := core.(domain.Core)
		res = EditUserResponseFormat{
			ID:       cnv.ID,
			Fullname: cnv.Fullname,
			Email:    cnv.Email,
			Location: cnv.Location,
			Profile:  cnv.Profile,
		}
	case "get":
		cnv := core.(domain.Core)
		res = GetUserResponseFormat{
			ID:       cnv.ID,
			Fullname: cnv.Fullname,
			Email:    cnv.Email,
			Location: cnv.Location,
			Profile:  cnv.Profile,
		}
	}
	return res
}
