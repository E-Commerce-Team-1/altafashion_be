package delivery

import "altafashion_be/feature/carts/domain"

type GetId struct {
	id uint `param:"id"`
}

type AddFormat struct {
	IdProduct uint `json:"id_product" form:"id_product"`
	UserID    uint `json:"user_id" form:"user_id"`
	Qty       int  `json:"qty" form:"qty"`
	Price     int  `json:"price" form:"price"`
	// Image     string `json:"image" form:"image"`
}

type UpdateFormat struct {
	ID uint `json:"id" form:"id"`
	//IdProduct uint `json:"id_product" form:"id_product"`
	UserID uint `json:"user_id" form:"user_id"`
	Qty    int  `json:"qty" form:"qty"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {

	case AddFormat:
		cnv := i.(AddFormat)
		return domain.Core{
			IdProduct: cnv.IdProduct,
			UserID:    cnv.UserID,
			Qty:       cnv.Qty,
			Price:     cnv.Price,
			// Image:     cnv.Image,
		}
	}
	return domain.Core{}
}
