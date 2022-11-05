package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/eqgo/eval"
)

func main() {
	ctx := eval.DefaultContext()
	switch len(os.Args) {
	case 1:
		doMultiMode(ctx)
	case 2:
		res, err := solveExpr(os.Args[1], ctx)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
	default:
		fmt.Println("too many arguments")
	}
}

func solveExpr(exprString string, ctx *eval.Context) (any, error) {
	expr := eval.New(exprString)
	err := expr.Compile(ctx)
	if err != nil {
		return nil, err
	}
	res, err := expr.Eval(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func doMultiMode(ctx *eval.Context) {
	fmt.Println()
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expr := scanner.Text()

	if expr == "exit" {
		return
	}

	res, err := solveExpr(expr, ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	doMultiMode(ctx)
}
