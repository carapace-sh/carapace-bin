package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/syft"
	"github.com/spf13/cobra"
)

var packagesCmd = &cobra.Command{
	Use:   "packages [SOURCE]",
	Short: "Generate a package SBOM",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packagesCmd).Standalone()
	addCommonFlags(packagesCmd)
	rootCmd.AddCommand(packagesCmd)

	carapace.Gen(packagesCmd).PositionalCompletion(
		syft.ActionSources(),
	)
}
