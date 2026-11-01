package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remote_help_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Summarize a remote (URLs, branches, pull config)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_help_showCmd).Standalone()

	remote_helpCmd.AddCommand(remote_help_showCmd)
}
