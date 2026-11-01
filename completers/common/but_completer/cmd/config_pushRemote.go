package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var config_pushRemoteCmd = &cobra.Command{
	Use:   "push-remote",
	Short: "View or set the remote used to push branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_pushRemoteCmd).Standalone()

	config_pushRemoteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_pushRemoteCmd)

	carapace.Gen(config_pushRemoteCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
