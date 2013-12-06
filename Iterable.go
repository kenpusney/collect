package collect

type Iterable interface {
	GetIterator() Iterator
}
