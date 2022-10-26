package delivery

import "altafashion_be/feature/carts/domain"

type GetId struct {
	id uint `param:"id"`
}

type AddFormat struct {
	IdProduct   uint   `json:"id_product" form:"id_product"`
	IdUser      uint   `json:"id_user" form:"id_user"`
	ProductName string `json:"product_name" form:"product_name"`
	ProductQty  uint   `json:"product_qty" form:"product_qty"`
	Price       int    `json:"price" form:"price"`
	Image       string `json:"image" form:"image"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {

	case AddFormat:
		cnv := i.(AddFormat)
		return domain.Core{
			IdProduct:   cnv.IdProduct,
			IdUser:      cnv.IdUser,
			ProductName: cnv.ProductName,
			ProductQty:  cnv.ProductQty,
			Price:       cnv.Price,
			Image:       cnv.Image,
		}
	}
	return domain.Core{}
}
