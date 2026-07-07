package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_managerCmd = &cobra.Command{
	Use:   "manager",
	Short: "Git credential helper backed by the Windows Credential Manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_managerCmd).Standalone()

	helpCmd.AddCommand(help_managerCmd)
}
