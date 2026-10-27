package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Summarize a remote (URLs, branches, pull config)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_showCmd).Standalone()

	remote_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	remote_showCmd.Flags().StringP("limit", "n", "", "Return at most this many remote branches")
	remote_showCmd.Flags().Bool("no-query", false, "Use local remote-tracking refs only; do not contact the remote")
	remoteCmd.AddCommand(remote_showCmd)

	carapace.Gen(remote_showCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
