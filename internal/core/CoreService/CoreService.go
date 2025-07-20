package CoreService

type Entity interface{ GetId() int }

type CoreService[T any] interface {
	Get() ([]T, error)
	FindById(id int) (T, error)
	Insert(entity T) (T, error)
	Update(entity T) (T, error)
	Delete(id int) error
}
