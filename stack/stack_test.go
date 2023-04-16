package stack

import (
	"reflect"
	"testing"
)

func TestStackPushPop(t *testing.T) {
	ts := NewTransactionStack()

	map1 := make(map[string]string)
	map1["a"] = "foo"
	map1["b"] = "bar"
	tx1 := Transaction{Store: map1}
	ts.Push(tx1)

	val := ts.Pop()
	if !reflect.DeepEqual(val, tx1) {
		t.Errorf("Expected %v, Got: %v", tx1, val)
	}
}
