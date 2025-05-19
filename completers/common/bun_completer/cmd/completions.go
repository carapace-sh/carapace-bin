package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionsCmd = &cobra.Command{
	Use:   "completions",
	Short: "Install shell completions for tab-completion",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionsCmd).Standalone()

	rootCmd.AddCommand(completionsCmd)

	carapace.Gen(completionsCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
