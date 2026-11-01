package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ignore_help_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Print whether paths match exclude rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ignore_help_checkCmd).Standalone()

	ignore_helpCmd.AddCommand(ignore_help_checkCmd)
}
