package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revisionCmd = &cobra.Command{
	Use:   "revision",
	Short: "Revision commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revisionCmd).Standalone()

	helpCmd.AddCommand(help_revisionCmd)
}
