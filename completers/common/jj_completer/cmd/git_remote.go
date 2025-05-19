package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage Git remotes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_remoteCmd).Standalone()

	git_remoteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gitCmd.AddCommand(git_remoteCmd)
}
