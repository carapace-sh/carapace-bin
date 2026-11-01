package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_remote_refsCmd = &cobra.Command{
	Use:   "refs",
	Short: "List references on a remote (like `git ls-remote`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_remote_refsCmd).Standalone()

	help_remoteCmd.AddCommand(help_remote_refsCmd)
}
