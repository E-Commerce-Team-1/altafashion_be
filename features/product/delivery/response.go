package delivery

import "altafashion_be/features/product/domain"

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

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessDeleteResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailedResponse(msg string) map[string]string {
	return map[string]string{
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
	}

	return res
}
