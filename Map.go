package collect

type Map interface {
	Insert(interface{}, interface{}) bool
	Find(interface{}) (interface{}, bool)
	Do(func(interface{}, interface{}))
	Len() int
}
