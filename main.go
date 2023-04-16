package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sowjumn/interview/devoted/DB"
)

func main() {
	mainDB := DB.NewDB()
	//transactionStack := stack.NewTransactionStack()

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
				SetVal(opArr, mainDB)
			case "GET":
				GetVal(opArr, mainDB)
			case "COUNT":
				CountKey(opArr, mainDB)
			case "DELETE":
				DeleteKey(opArr, mainDB)
			case "DEFAULT":
				PrintError()
			}
		}
	}
}

func SetVal(opArr []string, mainDB *DB.DB) {
	if len(opArr) != 3 {
		PrintError()
		os.Exit(1)
	}
	mainDB.Set(opArr[1], opArr[2])
}

func GetVal(opArr []string, mainDB *DB.DB) {
	if len(opArr) != 2 {
		PrintError()
		os.Exit(1)
	}
	val, ok := mainDB.Get(opArr[1])
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("NULL")
	}
}

func CountKey(opArr []string, mainDB *DB.DB) {
	if len(opArr) != 2 {
		PrintError()
		os.Exit(1)
	}
	val := mainDB.Count(opArr[1])
	fmt.Println(val)
}

func DeleteKey(opArr []string, mainDB *DB.DB) {
	if len(opArr) != 3 {
		PrintError()
		os.Exit(1)
	}
	mainDB.Delete(opArr[1])
}

func PrintError() {
	fmt.Println("Unrecognized operation")
}
