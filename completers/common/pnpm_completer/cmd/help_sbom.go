package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_sbomCmd = &cobra.Command{
	Use:   "sbom",
	Short: "Generate a Software Bill of Materials (SBOM)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_sbomCmd).Standalone()

	helpCmd.AddCommand(help_sbomCmd)
}
