package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_licensesCmd = &cobra.Command{
	Use:   "licenses",
	Short: "Check the licenses of the installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_licensesCmd).Standalone()

	helpCmd.AddCommand(help_licensesCmd)
}
