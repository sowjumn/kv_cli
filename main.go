package main

import (
	"github.com/sowjumn/interview/devoted/DB"
	"github.com/sowjumn/interview/devoted/stack"
)

func main() {
	mainDB := DB.NewDB()
	transactionStack := stack.NewTransactionStack()
}
