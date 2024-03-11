package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var git_remote_renameCmd = &cobra.Command{
	Use:   "rename [OPTIONS] <OLD> <NEW>",
	Short: "Rename a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_remote_renameCmd).Standalone()

	git_remote_renameCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_remoteCmd.AddCommand(git_remote_renameCmd)

	carapace.Gen(git_remote_renameCmd).PositionalCompletion(
		jj.ActionRemotes(),
		jj.ActionRemotes(),
	)
}
