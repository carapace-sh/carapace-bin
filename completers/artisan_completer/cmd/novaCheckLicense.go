package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaCheckLicenseCmd = &cobra.Command{
	Use:   "nova:check-license",
	Short: "Verify your Nova license key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaCheckLicenseCmd).Standalone()

	rootCmd.AddCommand(novaCheckLicenseCmd)
}
