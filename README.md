# solve
Solve is a simple command line utility to evaluate math expressions.

## Install

Use the following command to install solve:
    
    go install github.com/eqgo/solve@latest

The same command can also be used to update solve.

## Usage

You can directly solve an expression by using

    solve expr

The expression does not need to be in quotation marks unless it uses spaces or other symbols that interfere with the command line.

Example:

    solve 1+1
    2

Or:

    solve "3cospi"
    -3

You can also enter multi mode by calling solve with no arguments, which will then allow you to enter multiple expressions (also, no quotation marks are needed in multi mode):

    solve

    > 1-3
    -2

    > !true
    false

    > exit

You can exit multi mode by just typing `exit`.

## Expressions

Solve uses [eqgo/eval](https://github.com/eqgo/eval) to evaluate expressions, so it uses the same syntax, types, and operators as it. Solve uses `eval.DefaultContext()` for its context, so it supports all of the standard math functions, pi, and e.