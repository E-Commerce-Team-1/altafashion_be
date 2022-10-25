package delivery

import "altafashion_be/feature/users/domain"

type RegisterFormat struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}
type UpdateUserFormat struct {
	Fullname string `json:"fullname"  form:"fullname"`
	Email    string `json:"email"  form:"email"`
	Password string `json:"password"  form:"password"`
	Location string `json:"location"  form:"location"`
	Profile  string `json:"profile" form:"profile" `
}
type GetUserRequestFormat struct {
	Email string `json:"email"  form:"email"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Fullname: cnv.Fullname, Email: cnv.Email, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Email: cnv.Email, Password: cnv.Password}
	case UpdateUserFormat:
		cnv := i.(UpdateUserFormat)
		return domain.Core{
			Fullname: cnv.Fullname,
			Email:    cnv.Email,
			Password: cnv.Password,
			Location: cnv.Location,
		}
	case GetUserRequestFormat:
		cnv := i.(GetUserRequestFormat)
		return domain.Core{Email: cnv.Email}
	}
	return domain.Core{}
}
