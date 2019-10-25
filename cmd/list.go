package cmd

import(
	"github.com/mhallmark/gotodo/data"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use: "list",
	Short: "Lists todo items.",
	Long: "Lists todo items. Default filter is unfinished. Use -a to list all items.",
	Run: list,
}

func list(cmd *cobra.Command, args []string) {
	items, err := data.List()

	if (err != nil) {
		cmd.PrintErr(err)
		return
	}

	for item := range items {
		cmd.Println(item)
	}
}