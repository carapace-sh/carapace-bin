package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateLicenseDataCmd = &cobra.Command{
	Use:     "update-license-data",
	Short:   "Update SPDX license data in the Homebrew repository",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateLicenseDataCmd).Standalone()

	updateLicenseDataCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	updateLicenseDataCmd.Flags().BoolP("help", "h", false, "Show this message.")
	updateLicenseDataCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	updateLicenseDataCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(updateLicenseDataCmd)
}
