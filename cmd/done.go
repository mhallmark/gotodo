package cmd

import (
	"os"

	"github.com/mhallmark/gotodo/data"
	"github.com/spf13/cobra"
)

var doneCommand = &cobra.Command{
	Use:   "done",
	Short: "Toggles an item as done/undone",
	Long:  "Toggles an item as done/undone. Requires item ids.",
	Run:   done,
}

func done(cmd *cobra.Command, args []string) {
	changedItems, errs, done := data.Done(args)

	for {
		select {
		case <-done:
			return
		case err := <-errs:
			if err == nil {
				break
			}
			cmd.PrintErrln(err)
			os.ErrClosed = err
		case item := <-changedItems:
			cmd.Printf("Updated %v to DONE\n", item.ID.String()[0:8])
		}
	}
}
