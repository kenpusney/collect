package collect

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
