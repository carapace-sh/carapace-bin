package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_fsckCmd = &cobra.Command{
	Use:   "fsck",
	Short: "Check refs, objects, and pack indexes for problems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_fsckCmd).Standalone()

	helpCmd.AddCommand(help_fsckCmd)
}
