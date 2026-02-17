package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Prints the license of this software.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseCmd).Standalone()

	rootCmd.AddCommand(licenseCmd)
}
