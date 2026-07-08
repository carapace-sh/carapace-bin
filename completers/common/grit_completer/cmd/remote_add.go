package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_addCmd).Standalone()

	remote_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	remoteCmd.AddCommand(remote_addCmd)

	carapace.Gen(remote_addCmd).PositionalCompletion(
		carapace.ActionValues(),
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
	)
}
