package stack

import "github.com/sowjumn/interview/devoted/DB"

type Transaction struct {
	Cache *DB.DB
}

func NewTransaction() *Transaction {
	return &Transaction{
		Cache: DB.NewDB(),
	}
}

type TransactionStack struct {
	Stack []Transaction
}

func (ts *TransactionStack) Push(t Transaction) {
	ts.Stack = append(ts.Stack, t)
}

func (ts *TransactionStack) Pop() Transaction {
	if len(ts.Stack) == 0 {
		return Transaction{}
	}
	n := len(ts.Stack) - 1
	lastTx := ts.Stack[n]
	ts.Stack = ts.Stack[:n]
	return lastTx
}

func (ts *TransactionStack) Empty() {
	ts.Stack = []Transaction{}
}

func NewTransactionStack() *TransactionStack {
	return &TransactionStack{
		Stack: []Transaction{},
	}
}
