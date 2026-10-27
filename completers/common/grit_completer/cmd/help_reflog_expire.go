package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_reflog_expireCmd = &cobra.Command{
	Use:   "expire",
	Short: "Drop reflog entries older than configured or explicit cutoffs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_reflog_expireCmd).Standalone()

	help_reflogCmd.AddCommand(help_reflog_expireCmd)
}
