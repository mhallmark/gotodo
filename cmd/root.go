package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "A todo list cli.",
		Long: `A todo list cli.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addCommand)
	rootCmd.AddCommand(listCommand)
	rootCmd.AddCommand(removeCommand)
}