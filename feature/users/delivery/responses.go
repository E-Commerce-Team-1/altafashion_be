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
	Username string `json:"username"`
	Email    string `json:"email"`
}

type loginResponses struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type EditUserResponseFormat struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}
type GetUserResponseFormat struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}

func ToResponse(core interface{}, code string, token string) interface{} {
	var res interface{}

	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = registerRespons{ID: cnv.ID, Username: cnv.Username, Email: cnv.Email}
	case "login":
		cnv := core.(domain.Core)
		res = loginResponses{Username: cnv.Username, Email: cnv.Email, Token: token}
	case "edit":
		cnv := core.(domain.Core)
		res = EditUserResponseFormat{
			ID:       cnv.ID,
			Username: cnv.Username,
			Email:    cnv.Email,
			Phone:    cnv.Phone,
			Location: cnv.Location,
		}
	case "get":
		cnv := core.(domain.Core)
		res = GetUserResponseFormat{
			ID:       cnv.ID,
			Username: cnv.Username,
			Email:    cnv.Email,
			Phone:    cnv.Phone,
			Location: cnv.Location,
		}
	}
	return res
}
