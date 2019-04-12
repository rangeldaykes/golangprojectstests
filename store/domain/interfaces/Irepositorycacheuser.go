package interfaces

type IRepositoryCacheUser interface {
	GetUser(ID string) (string, error)
}
