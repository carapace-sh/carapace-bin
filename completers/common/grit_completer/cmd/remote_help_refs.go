package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remote_help_refsCmd = &cobra.Command{
	Use:   "refs",
	Short: "List references on a remote (like `git ls-remote`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_help_refsCmd).Standalone()

	remote_helpCmd.AddCommand(remote_help_refsCmd)
}
