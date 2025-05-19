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

	updateLicenseDataCmd.Flags().Bool("debug", false, "Display any debugging information.")
	updateLicenseDataCmd.Flags().Bool("help", false, "Show this message.")
	updateLicenseDataCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	updateLicenseDataCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(updateLicenseDataCmd)
}
