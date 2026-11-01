package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reflog_help_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove one or more reflog entries (`ref@{n}`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reflog_help_deleteCmd).Standalone()

	reflog_helpCmd.AddCommand(reflog_help_deleteCmd)
}
