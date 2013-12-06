package collect

import (
	"strings"
)

func CaseFoldedComparor(a, b interface{}) bool {
	return strings.ToLower(a.(string)) < strings.ToLower(b.(string))
}

func IntComparor(a, b interface{}) bool {
	return a.(int) < b.(int)
}

func StringComparor(a, b interface{}) bool {
	return a.(string) < b.(string)
}
