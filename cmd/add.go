package cmd

import (
	"github.com/mhallmark/gotodo/data/todoitems"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a todo item.",
	Long:  "Add a todo item.",
	Run:   add,
}

func add(cmd *cobra.Command, args []string) {
	if len(args) <= 0 {
		cmd.PrintErr("Must provide at least one string to set as a todo item.\n")
		cmd.ErrOrStderr()
		return
	}
	
	todoItems, errs, done := todoitems.Add(args)
	
	for {
		select {
		case i := <- done:
			cmd.Printf("Added %v total items.\n", i)
			return
		case item := <- todoItems:
			cmd.Printf("Added \"%v\".\n", item.Message)
		case err := <- errs:
			cmd.PrintErrln(err)
		}
	}
}
