package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_absorbCmd = &cobra.Command{
	Use:   "absorb",
	Short: "Amends changes into the appropriate commits where they belong",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_absorbCmd).Standalone()

	helpCmd.AddCommand(help_absorbCmd)
}
