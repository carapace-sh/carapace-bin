package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open [url]",
	Short: "Open the browser in a terminal pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openCmd).Standalone()

	openCmd.Flags().Bool("allow-clipboard-read", false, "Lets websites read from clipboard")
	openCmd.Flags().Bool("no-merge", false, "Do not open as a tab in a neighbor terminal-browser")
	openCmd.Flags().String("size", "", "How much of the space the split takes (0.2 to 0.95)")
	openCmd.Flags().String("split", "", "Open in a new pane")
	openCmd.Flags().String("ssh", "", "Perform all network requests through a remote server")
	rootCmd.AddCommand(openCmd)

	carapace.Gen(openCmd).FlagCompletion(carapace.ActionMap{
		"split": carapace.ActionValues("right", "left", "down", "up"),
	})

	carapace.Gen(openCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
