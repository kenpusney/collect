package test

import (
	"collect"
	"testing"
)

func Test_IntComparator(t *testing.T) {
	if collect.IntComparor(1, 2) {
		t.Log("success!")
	} else {
		t.Error("Comparion failure!")
	}
}

func Test_StringComparator(t *testing.T) {
	if collect.StringComparor("abc", "def") {
		t.Log("success!")
	} else {
		t.Error("Comparion failure!")
	}
}

func Test_CaseFoldedComparator(t *testing.T) {
	if collect.CaseFoldedComparor("abc", "DEF") {
		t.Log("success!")
	} else {
		t.Error("Comparion failure!")
	}
}
