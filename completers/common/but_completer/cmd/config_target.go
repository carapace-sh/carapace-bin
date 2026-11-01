package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var config_targetCmd = &cobra.Command{
	Use:   "target",
	Short: "View or set the target branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_targetCmd).Standalone()

	config_targetCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_targetCmd.Flags().String("push-remote", "", "Remote to push branches to (e.g., \"origin\" for a fork)")
	configCmd.AddCommand(config_targetCmd)

	carapace.Gen(config_targetCmd).FlagCompletion(carapace.ActionMap{
		"push-remote": git.ActionRemotes(),
	})

	carapace.Gen(config_targetCmd).PositionalCompletion(
		git.ActionRemoteBranches(""), // TODO local branches as well?
	)
}
