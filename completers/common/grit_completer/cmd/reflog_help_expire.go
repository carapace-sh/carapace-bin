package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reflog_help_expireCmd = &cobra.Command{
	Use:   "expire",
	Short: "Drop reflog entries older than configured or explicit cutoffs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reflog_help_expireCmd).Standalone()

	reflog_helpCmd.AddCommand(reflog_help_expireCmd)
}
