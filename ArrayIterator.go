// ArrayIterator is a wrapper for native array
package collect

type ArrayIterator struct {
	array   []interface{}
	length  int
	current int
}

func (i *ArrayIterator) HasNext() bool {
	return i.current < i.length
}

func (a *ArrayIterator) Next() interface{} {
	value := a.array[a.current]
	a.current++
	return value
}

func (i *ArrayIterator) GetIterator() Iterator {
	return i
}

func NewArrayIterator(array []interface{}) *ArrayIterator {
	return &ArrayIterator{array: array, length: len(array)}
}
