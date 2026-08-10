package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_computeChecksumCmd = &cobra.Command{
	Use:   "compute-checksum",
	Short: "Compute the checksum for a binary artifact",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_computeChecksumCmd).Standalone()
	package_computeChecksumCmd.Flags().SetInterspersed(false)

	package_computeChecksumCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_computeChecksumCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_computeChecksumCmd)

	carapace.Gen(package_computeChecksumCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
