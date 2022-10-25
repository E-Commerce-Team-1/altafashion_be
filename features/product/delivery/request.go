package delivery

import "altafashion_be/features/product/domain"

type ProductFormat struct {
	ID          uint   `json:"id" form:"id"`
	Image       string `json:"image" form:"image"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Category    string `json:"category" form:"category"`
	Qty         int    `josn:"qty" form:"qty"`
	Price       int    `json:"price" form:"price"`
	UserID      uint   `json:"id_user" form:"id_user"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case ProductFormat:
		cnv := i.(ProductFormat)
		return domain.Core{
			ID:          cnv.ID,
			Image:       cnv.Image,
			Name:        cnv.Name,
			Description: cnv.Description,
			Category:    cnv.Category,
			Qty:         cnv.Qty,
			Price:       cnv.Price,
			UserID:      cnv.UserID,
		}
	}

	return domain.Core{}
}
