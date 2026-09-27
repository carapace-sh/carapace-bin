package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_whyCmd = &cobra.Command{
	Use:   "why",
	Short: "Shows the packages that depend on `pkg`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_whyCmd).Standalone()

	helpCmd.AddCommand(help_whyCmd)
}
