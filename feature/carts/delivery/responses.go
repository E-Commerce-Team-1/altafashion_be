package delivery

import "altafashion_be/feature/carts/domain"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

type AddResponse struct {
	IdProduct  uint `json:"id_product"`
	ProductQty uint `json:"product_qty"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "register":
		cnv := core.(domain.Core)
		res = AddResponse{IdProduct: cnv.IdProduct, ProductQty: cnv.ProductQty}
		// case "update":
		// 	cnv := core.(domain.Core)
		// 	res = UpdateResponse{ID: cnv.ID, ProductQty: cnv.ProductQty}
		//
	}

	return res
}
