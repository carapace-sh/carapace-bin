package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var newTabCmd = &cobra.Command{
	Use:   "new-tab [url]",
	Short: "Open a tab here, and a browser too if there is none",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newTabCmd).Standalone()

	newTabCmd.Flags().String("browser", "", "A browser key from terminal-browser ls")
	rootCmd.AddCommand(newTabCmd)

	carapace.Gen(newTabCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
