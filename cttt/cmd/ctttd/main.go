package main

import (
	"os"

	"github.com/cosmos_tic_tac_toe/cttt/cmd/ctttd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
