package cmd

import(
	"github.com/mhallmark/gotodo/data"
	"github.com/spf13/cobra"
)

var removeCommand = &cobra.Command{
	Use: "remove",
	Short: "Removes an item.",
	Long: "Removes an item by id.",
	Run: remove,
}

func remove(cmd *cobra.Command, args []string) {
	for _, key := range args {
		var err = data.Remove(key)
		if (err != nil) {
			cmd.PrintErr(err)
			continue
		}
		cmd.Printf("Removed %v.\n", key)
	}
}