package user

type IRepository interface {
	Save(user *Entity) error
	FindByUsername(username string) (userEntity *Entity, err error)
}
