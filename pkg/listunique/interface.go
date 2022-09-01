package listunique

type Interface[T comparable] interface {
	Add(elem T)
	AddAll(elems ...T)
	GetList() []T
	GetSize() int
	Update(index int, elem T) bool
	Delete(elem T) bool
}
