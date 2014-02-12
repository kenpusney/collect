package util

import (
	"collect"
	//"collect/algo"
)

func Listify(it collect.Iterator) collect.List {
	list := &collect.LinkedList{}
	for it.HasNext() {
		list.Append(it.Next())
	}
	return list
	//return (algo.Reduce(it, func(list *interface{}, e interface{}) {
	//	((*list).(collect.List)).Append(e)
	//}, &collect.LinkedList{})).(collect.List)
}
