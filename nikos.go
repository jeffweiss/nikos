package main

import (
	"github.com/jeffweiss/nikos/cmd"
)

func main() {
	cmd.SetupCommands()
	cmd.RootCmd.Execute()
}
