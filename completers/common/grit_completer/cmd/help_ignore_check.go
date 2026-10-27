package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ignore_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Print whether paths match exclude rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ignore_checkCmd).Standalone()

	help_ignoreCmd.AddCommand(help_ignore_checkCmd)
}
