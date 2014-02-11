package algo

import (
	"collect"
)

func ForEach(iter collect.Iterator, fun func(interface{})) {
	for iter.HasNext() {
		fun(iter.Next())
	}
}

func Reduce(iter collect.Iterator,
	fun func(*interface{}, interface{}),
	init interface{}) interface{} {
	for iter.HasNext() {
		fun(&init, iter.Next())
	}
	return init
}

func Map(iter collect.Iterator,
	fun func(interface{}) interface{}) collect.Iterator {
	return (Reduce(iter, func(list *interface{}, item interface{}) {
		((*list).(collect.List)).Append(fun(item))
	}, &collect.LinkedList{})).(collect.List).GetIterator()
}
