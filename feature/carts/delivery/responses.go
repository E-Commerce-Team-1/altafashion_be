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
	IdProduct uint `json:"id_product"`
	Qty       int  `json:"qty"`
}

func SuccessResponseNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

type GetProd struct {
	ID        uint   `json:"id"`
	IdProduct uint   `json:"id_product"`
	UserID    uint   `json:"user_id"`
	NamaToko  string `json:"nama_toko"`
	Name      string `json:"name"`
	Qty       int    `json:"qty"`
	Price     int    `json:"price"`
	Image     string `json:"image"`
}

type UpdateResponse struct {
	ID  uint `json:"id"`
	Qty int  `json:"qty"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "add":
		cnv := core.(domain.Core)
		res = AddResponse{IdProduct: cnv.IdProduct, Qty: cnv.Qty}
	case "update":
		cnv := core.(domain.Core)
		res = UpdateResponse{ID: cnv.ID, Qty: cnv.Qty}
	}
	return res
}

func ToResponseProduct(core interface{}, code string) interface{} {
	var res interface{}
	var arr []GetProd
	val := core.([]domain.Core)
	for _, cnv := range val {
		arr = append(arr, GetProd{ID: cnv.ID, IdProduct: cnv.IdProduct, UserID: cnv.UserID, Name: cnv.Name,
			Qty: cnv.Qty, Price: cnv.Price, Image: cnv.Image})
	}
	res = arr
	return res
}
