package cmd

import (
	"time"
	"github.com/mhallmark/gotodo/data"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use: "add",
	Short: "Add a todo item.",
	Long: "Add a todo item.",
	Run: add,
}

func add(cmd *cobra.Command, args []string) {
	if len(args) <= 0 {
		cmd.PrintErr("Must provide at least one string to set as a todo item.\n");
		cmd.ErrOrStderr()
		return
	}

	for _, v := range args {
		id, err := uuid.NewUUID();

		if err != nil {
			cmd.PrintErr("Oops, something done goofed.")
			cmd.ErrOrStderr();
			return;
		}

		todo := data.TodoItem {
			Id: id,
			Message: v,
			Created: time.Now(),
			Done: false,
		}

		aErr := data.Add(todo)

		if (aErr != nil) {
			cmd.PrintErr("Oops, something done goofed.")
			cmd.ErrOrStderr();
			return;
		}

		cmd.Printf("Added \"%v\".\n", todo.Message) 
	}

	cmd.Printf("Added %v new items.\n", len(args))
}