package collect

type Iterator interface {
	Iterable
	HasNext() bool
	Next() interface{}
}
