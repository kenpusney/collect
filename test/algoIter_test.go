package test

import (
	"collect"
	"collect/algo"
	"fmt"
	"testing"
)

func Test_Reduce(t *testing.T) {
	var list collect.List = &collect.LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Put(2, 5)
	itor := list.GetIterator()
	algo.ForEach(itor, func(v interface{}) {
		fmt.Println(v.(int))
	})
	var d = 0
	fmt.Println(algo.Reduce(list.GetIterator(), func(v *interface{}, e interface{}) {
		(*v) = (*v).(int) + e.(int)
	}, d))
	algo.ForEach(algo.Map(list.GetIterator(), func(v interface{}) interface{} {
		return (v).(int) * (v).(int)
	}), func(v interface{}) {
		fmt.Println(v.(int))
	})
	t.Log("Success!")
}
