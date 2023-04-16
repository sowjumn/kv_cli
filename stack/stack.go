package stack

type Transaction struct {
	Store map[string]string
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

func NewTransactionStack() TransactionStack {
	return TransactionStack{
		Stack: []Transaction{},
	}
}
