package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "Analyze current tools and check if newer versions are available",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()

	outdatedCmd.Flags().Bool("bump", false, "Bump versions in mise.toml")
	rootCmd.AddCommand(outdatedCmd)

	carapace.Gen(outdatedCmd).PositionalAnyCompletion(
		action.ActionTools(),
	)
}
