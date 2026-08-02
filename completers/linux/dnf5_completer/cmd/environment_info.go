package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environmentInfoCmd = &cobra.Command{
	Use:   "info [options] [<environment-spec>...]",
	Short: "print details about comps environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environmentInfoCmd).Standalone()

	environmentInfoCmd.Flags().Bool("available", false, "List available environments")
	environmentInfoCmd.Flags().Bool("installed", false, "List installed environments")

	environmentCmd.AddCommand(environmentInfoCmd)
}
