package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach to a session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_attachCmd).Standalone()

	helpCmd.AddCommand(help_attachCmd)
}
