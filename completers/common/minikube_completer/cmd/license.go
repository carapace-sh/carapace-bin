package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Outputs the licenses of dependencies to a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseCmd).Standalone()

	licenseCmd.Flags().StringP("dir", "d", "", "Directory to output licenses to")
	rootCmd.AddCommand(licenseCmd)

	carapace.Gen(licenseCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
	})
}
