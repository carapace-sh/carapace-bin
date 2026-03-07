package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var hyperlinkedGrepCmd = &cobra.Command{
	Use:   "hyperlinked-grep",
	Short: "Add hyperlinks to the output of ripgrep",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hyperlinkedGrepCmd).Standalone()

	rootCmd.AddCommand(hyperlinkedGrepCmd)

	carapace.Gen(hyperlinkedGrepCmd).PositionalAnyCompletion(
		bridge.ActionCarapace("ssh"),
	)
}
