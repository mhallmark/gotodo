package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version for gotodo.",
  Long:  `Print the version for gotodo.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("gotodo - version 1.0")
  },
}