package main

import (
	cmd "todo/cmd/todo"
	"todo/internal/utils/must"
)

func main() {
	must.Panic(cmd.Root.Execute())
}
