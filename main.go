package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sowjumn/interview/devoted/DB"
	"github.com/sowjumn/interview/devoted/stack"
)

func main() {
	mainDB := DB.NewDB()
	txStack := stack.NewTransactionStack()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(">> ")
		op, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(0)
		}
		opArr := strings.Fields(op)
		if len(opArr) > 0 {
			switch opArr[0] {
			case "END":
				os.Exit(0)
			case "SET":
				SetVal(opArr, mainDB, txStack)
			case "GET":
				GetVal(opArr, mainDB, txStack)
			case "COUNT":
				CountKey(opArr, mainDB, txStack)
			case "DELETE":
				DeleteKey(opArr, mainDB, txStack)
			case "BEGIN":
				StartTransaction(txStack)
			case "COMMIT":
				CommitTransaction(mainDB, txStack)
			case "ROLLBACK":
				RollbackTransaction(txStack)
			case "DEFAULT":
				PrintError()
			}
		}
	}
}

func SetVal(opArr []string, mainDB *DB.DB, currTxStack *stack.TransactionStack) {
	if len(opArr) != 3 {
		PrintError()
		os.Exit(1)
	}

	if len(currTxStack.Stack) == 0 {
		// if transaction stack empty
		mainDB.Set(opArr[1], opArr[2])
	} else {
		// get the last transaction
		recentTx := currTxStack.Pop()
		recentTx.Cache.Set(opArr[1], opArr[2])
		currTxStack.Push(recentTx)
	}
}

func GetVal(opArr []string, mainDB *DB.DB, currTxStack *stack.TransactionStack) {
	var val string
	var ok bool

	if len(opArr) != 2 {
		PrintError()
		os.Exit(1)
	}

	// walk through the stack implemented as a slice of txs to find the latest value for the key
	for i := len(currTxStack.Stack) - 1; i >= 0; i-- {
		lastTx := currTxStack.Stack[i]
		val, ok = lastTx.Cache.Get(opArr[1])
		if ok {
			break
		}
	}

	// if key is not found in the transaction stack get it from mainDB
	if !ok {
		val, ok = mainDB.Get(opArr[1])
	}

	if ok && val != "DELETE" {
		fmt.Println(val)
	} else {
		fmt.Println("NULL")
	}

}

func CountKey(opArr []string, mainDB *DB.DB, txStack *stack.TransactionStack) {
	if len(opArr) != 2 {
		PrintError()
		os.Exit(1)
	}

	var currTx stack.Transaction
	setUnion := make(map[string]bool)
	count := mainDB.Count(opArr[1])
	for i := 0; i < len(txStack.Stack); i++ {
		currTx = txStack.Stack[i]
		// get the set union of the keys
		for k, v := range currTx.Cache.Store {
			if v == "DELETE" {
				if setUnion[k] {
					delete(setUnion, k)
				}
			} else if v == opArr[1] {
				setUnion[k] = true
			}
		}
	}

	count += len(setUnion)
	fmt.Println(count)
}

func DeleteKey(opArr []string, mainDB *DB.DB, currTxStack *stack.TransactionStack) {
	if len(opArr) != 2 {
		PrintError()
		os.Exit(1)
	}

	if len(currTxStack.Stack) > 0 {
		recentTx := currTxStack.Pop()
		recentTx.Cache.Set(opArr[1], "DELETE")
		currTxStack.Push(recentTx)
	} else {
		mainDB.Delete(opArr[1])
	}
	mainDB.Delete(opArr[1])
}

func StartTransaction(txStack *stack.TransactionStack) {
	newTx := stack.NewTransaction()
	txStack.Push(*newTx)
}

func CommitTransaction(mainDB *DB.DB, txStack *stack.TransactionStack) {
	var currTx stack.Transaction
	for i := 0; i < len(txStack.Stack); i++ {
		currTx = txStack.Stack[i]
		for k, v := range currTx.Cache.Store {
			if v == "DELETE" {
				mainDB.Delete(k)
			} else {
				mainDB.Set(k, v)
			}
		}
	}
	txStack.Empty()
}

func RollbackTransaction(txStack *stack.TransactionStack) {
	if len(txStack.Stack) > 0 {
		txStack.Pop()
	} else {
		fmt.Printf("TRANSACTION NOT FOUND")
	}
}

func PrintError() {
	fmt.Println("Unrecognized operation")
}
