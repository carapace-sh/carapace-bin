package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_remote_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Summarize a remote (URLs, branches, pull config)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_remote_showCmd).Standalone()

	help_remoteCmd.AddCommand(help_remote_showCmd)
}
