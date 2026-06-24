package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scalarCmd = &cobra.Command{
	Use:     "scalar",
	Short:   "A tool for managing large Git repositories",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(scalarCmd).Standalone()

	rootCmd.AddCommand(scalarCmd)

	carapace.Gen(scalarCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)
}
