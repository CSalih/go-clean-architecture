package main

import (
	"fmt"
	"os"

	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/command"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: "v0.1.0",
}

func init() {
	rootCmd.AddCommand(command.UserCommand)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
