package cmd

import (
	"github.com/mhallmark/gotodo/data"
	"github.com/spf13/cobra"
)

var removeCommand = &cobra.Command{
	Use:   "remove",
	Short: "Removes an item.",
	Long:  "Removes an item by id.",
	Run:   remove,
}

func remove(cmd *cobra.Command, args []string) {
	keys, errs, done := data.Remove(args)

	for {
		select {
		case <-done:
			return
		case key := <-keys:
			cmd.Printf("Removed %v.\n", key)
		case err := <-errs:
			cmd.PrintErrln(err)
		}
	}
}
