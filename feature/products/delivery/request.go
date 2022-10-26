package delivery

import "altafashion_be/feature/products/domain"

type AddProductFormat struct {
	ID          uint   `json:"id" form:"id"`
	Image       string `json:"image" form:"image"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Category    string `json:"category" form:"category"`
	Qty         uint   `josn:"qty" form:"qty"`
	Price       int    `json:"price" form:"price"`
	UserID      uint   `json:"id_user" form:"id_user"`
}

type EditProductFormat struct {
	ID          uint   `json:"id" form:"id"`
	Image       string `json:"image" form:"image"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Category    string `json:"category" form:"category"`
	Qty         uint   `josn:"qty" form:"qty"`
	Price       int    `json:"price" form:"price"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case AddProductFormat:
		cnv := i.(AddProductFormat)
		return domain.Core{
			ID:          cnv.ID,
			Image:       cnv.Image,
			Name:        cnv.Name,
			Description: cnv.Description,
			Category:    cnv.Category,
			Qty:         cnv.Qty,
			Price:       cnv.Price,
			// UserID:      cnv.UserID,
		}
	case EditProductFormat:
		cnv := i.(AddProductFormat)
		return domain.Core{
			ID:          cnv.ID,
			Image:       cnv.Image,
			Name:        cnv.Name,
			Description: cnv.Description,
			Category:    cnv.Category,
			Qty:         cnv.Qty,
			Price:       cnv.Price,
		}
	}

	return domain.Core{}
}
