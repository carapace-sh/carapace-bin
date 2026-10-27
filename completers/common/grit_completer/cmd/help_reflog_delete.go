package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_reflog_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove one or more reflog entries (`ref@{n}`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_reflog_deleteCmd).Standalone()

	help_reflogCmd.AddCommand(help_reflog_deleteCmd)
}
