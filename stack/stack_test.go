package stack

import (
	"reflect"
	"testing"

	"github.com/sowjumn/interview/devoted/DB"
)

func TestStackPushPop(t *testing.T) {
	ts := NewTransactionStack()

	localStore := DB.NewDB()
	tx1 := Transaction{Cache: localStore}
	ts.Push(tx1)

	val := ts.Pop()
	if !reflect.DeepEqual(val, tx1) {
		t.Errorf("Expected %v, Got: %v", tx1, val)
	}
}
