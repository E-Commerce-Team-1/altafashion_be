package domain

type Core struct {
	ID          uint
	Image       string
	Name        string
	Description string
	Category    string
	Qty         uint
	Price       int
	UserID      uint
}

type Repository interface {
	ShowAll(category, name string, page int) ([]Core, error)
	Insert(newProduct Core) (Core, error)
	ShowDetail(ID uint) (Core, error)
	Update(updateData Core, ID uint) (Core, error)
	Delete(ID uint) error
	ShowMyProduct() ([]Core, error)
}

type Services interface {
	GetAll(category, name string, page int) ([]Core, error)
	AddProduct(newProduct Core) (Core, error)
	GetDetail(ID uint) (Core, error)
	EditProduct(updateData Core, ID uint) (Core, error)
	Destroy(ID uint) error
	GetMyProduct() ([]Core, error)
}
