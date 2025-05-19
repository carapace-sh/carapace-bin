package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_remote_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Git remotes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_remote_listCmd).Standalone()

	git_remote_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_remoteCmd.AddCommand(git_remote_listCmd)
}
