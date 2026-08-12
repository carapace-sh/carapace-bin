package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_revert_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve conflicts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_revert_resolveCmd).Standalone()

	help_revision_revertCmd.AddCommand(help_revision_revert_resolveCmd)
}
