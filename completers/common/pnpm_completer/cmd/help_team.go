package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Manage organization teams and team memberships",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_teamCmd).Standalone()

	helpCmd.AddCommand(help_teamCmd)
}
