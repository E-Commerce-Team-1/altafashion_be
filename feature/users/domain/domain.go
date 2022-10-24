package domain

type Core struct {
	ID       uint
	Username string
	Email    string
	Password string
	Profile  string
	Phone    string
	Location string
}

type Repository interface {
	Insert(newUser Core) (Core, error)
	GetUser(existUser Core) (Core, error)
	Update(updateData Core, id uint) (Core, error)
	Delete(id uint) (Core, error)
	GetByUsername(username string) (Core, error)
	GetMyUser(id uint) (Core, error)
}

type Service interface {
	Register(newUser Core) (Core, error)
	Login(existUser Core) (Core, string, error)
	UpdateProfile(updateData Core, id uint) (Core, error)
	Deactivate(id uint) (Core, error)
	ShowByUsername(username string) (Core, error)
	MyProfile(id uint) (Core, error)
}
