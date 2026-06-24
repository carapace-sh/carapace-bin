package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scalar_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Scalar version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scalar_versionCmd).Standalone()

	scalar_versionCmd.Flags().Bool("build-options", false, "also print build options")
	scalar_versionCmd.Flags().BoolP("verbose", "v", false, "include Git version")
	scalarCmd.AddCommand(scalar_versionCmd)
}
