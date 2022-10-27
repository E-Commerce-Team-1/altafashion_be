package delivery

import "altafashion_be/feature/products/domain"

type ProductResponse struct {
	ID          uint   `json:"id" form:"id"`
	Image       string `json:"image" form:"image"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Category    string `json:"category" form:"category"`
	Qty         int    `json:"qty" form:"qty"`
	Price       int    `json:"price" form:"price"`
	UserID      uint   `json:"id_user" form:"id_user"`
}

type GetAllResponse struct {
	ID    uint   `json:"id" form:"id"`
	Image string `json:"image" form:"image"`
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessNoDataResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailedResponse(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "add":
		cnv := core.(domain.Core)
		res = ProductResponse{
			ID: cnv.ID, Image: cnv.Image, Name: cnv.Name, Description: cnv.Description,
			Category: cnv.Category, Qty: cnv.Qty, Price: cnv.Price, UserID: cnv.UserID,
		}
	case "edit":
		cnv := core.(domain.Core)
		res = ProductResponse{
			ID: cnv.ID, Image: cnv.Image, Name: cnv.Name, Description: cnv.Description,
			Category: cnv.Category, Qty: cnv.Qty, Price: cnv.Price,
		}
	}

	return res
}

func ToResponseList(core interface{}) interface{} {
	var res interface{}
	var list []GetAllResponse
	val := core.([]domain.Core)
	for _, cnv := range val {
		list = append(list, GetAllResponse{
			ID:    cnv.ID,
			Image: cnv.Image,
			Name:  cnv.Name,
			Price: cnv.Price,
		})
	}
	res = list

	return res
}
