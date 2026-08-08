package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environmentListCmd = &cobra.Command{
	Use:   "list [options] [<environment-spec>...]",
	Short: "list comps environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environmentListCmd).Standalone()

	environmentListCmd.Flags().Bool("available", false, "List available environments")
	environmentListCmd.Flags().Bool("installed", false, "List installed environments")

	environmentCmd.AddCommand(environmentListCmd)
}
