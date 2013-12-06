package collect

type List interface {
	Iterable
	Append(interface{}) bool
	Get(int) interface{}
	Put(int, interface{}) bool
}
