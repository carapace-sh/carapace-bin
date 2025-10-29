package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Modify a toolchain's installed components",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_componentCmd).Standalone()

	helpCmd.AddCommand(help_componentCmd)
}
