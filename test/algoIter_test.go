package test

import (
	"github.com/kenpusney/collect"
	"github.com/kenpusney/collect/algo"
	"testing"
)

func Test_algo_Reduce(t *testing.T) {
	var list collect.List = &collect.LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	sum := algo.Reduce(list.GetIterator(), func(v *interface{}, e interface{}) {
		(*v) = (*v).(int) + e.(int)
	}, 0)
	if sum == 6 {
		t.Log("Success!")
	} else {
		t.Error("Failed to reduce")
	}
}

func Test_algo_Map(t *testing.T) {
	var list collect.List = &collect.LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	sum := algo.Reduce(algo.Map(list.GetIterator(), func(v interface{}) interface{} {
		return (v).(int) * (v).(int)
	}), func(v *interface{}, e interface{}) {
		(*v) = (*v).(int) + e.(int)
	}, 0)
	if sum == 14 {
		t.Log("success!")
	} else {
		t.Errorf("Failed to Map %d", sum)
	}
}

func Test_algo_ForEach(t *testing.T) {
	var list collect.List = &collect.LinkedList{}
	count := 0
	list.Append(1)
	list.Append(2)
	list.Append(3)
	algo.ForEach(list.GetIterator(), func(v interface{}) {
		if v.(int) > 0 && v.(int) < 4 {
			count = count + 1
		}
	})
	if count == 3 {
		t.Log("Success!")
	} else {
		t.Error("Failed to ForEach")
	}
}

func Benchmark_algo_Map(b *testing.B) {
	b.StopTimer()
	var list collect.List = &collect.LinkedList{}
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		algo.Map(list.GetIterator(), func(v interface{}) interface{} {
			return v
		})
	}
}

func Benchmark_algo_ForEach(b *testing.B) {
	b.StopTimer()
	var list collect.List = &collect.LinkedList{}
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		algo.ForEach(list.GetIterator(), func(v interface{}) {

		})
	}
}

func Benchmark_algo_Reduce(b *testing.B) {
	b.StopTimer()
	var list collect.List = &collect.LinkedList{}
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		algo.Reduce(list.GetIterator(), func(v *interface{}, e interface{}) {
			(*v) = (*v).(int) + (e).(int)
		}, 0)
	}
}
