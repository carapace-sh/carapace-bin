package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/fury"
	"github.com/spf13/cobra"
)

var versionsCmd = &cobra.Command{
	Use:   "versions PACKAGE",
	Short: "List versions for a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionsCmd).Standalone()

	rootCmd.AddCommand(versionsCmd)

	carapace.Gen(versionsCmd).PositionalCompletion(
		fury.ActionPackages().MultiParts(":"),
	)
}
