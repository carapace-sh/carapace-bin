package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_overrideCmd = &cobra.Command{
	Use:   "override",
	Short: "Modify toolchain overrides for directories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_overrideCmd).Standalone()

	helpCmd.AddCommand(help_overrideCmd)
}
