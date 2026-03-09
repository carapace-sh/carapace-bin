package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environmentCmd = &cobra.Command{
	Use:   "environment <subcommand> [options] [<environment-spec>...]",
	Short: "manage comps environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environmentCmd).Standalone()

	rootCmd.AddCommand(environmentCmd)
}
