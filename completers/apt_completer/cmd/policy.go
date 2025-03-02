package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var policyCmd = &cobra.Command{
	Use:   "policy [pattern]...",
	Short: "show package priorities",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policyCmd).Standalone()

	rootCmd.AddCommand(policyCmd)

	carapace.Gen(policyCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch().FilterArgs(),
	)
}
