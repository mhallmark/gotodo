package cmd

import (
	"time"

	"github.com/mhallmark/gotodo/data/todoitems"
	"github.com/spf13/cobra"
)

var (
	allItems    bool
	listCommand = &cobra.Command{
		Use:     "list",
		Short:   "Lists todo items.",
		Long:    "Lists todo items. Default filter is unfinished. Use -a to list all items.",
		Example: "gotodo list",
		Run:     list,
	}
)

func init() {
	listCommand.Flags().BoolVarP(&allItems, "all", "a", false, "gotodo list --all true")
}

func list(cmd *cobra.Command, args []string) {
	items, errs, done := todoitems.List(allItems)

	for {
		select {
		case <-done:
			return
		case item := <-items:
			if allItems {
				cmd.Printf("%v \"%v\" DONE:%v\n", string(item.ID.String()[0:8]), item.Message, item.Done)
			} else {
				cmd.Printf("%v \"%v\" %v\n", string(item.ID.String()[0:8]), item.Message, item.Created.Format(time.UnixDate))
			}
		case err := <-errs:
			cmd.PrintErrln(err)
			return
		}
	}
}
