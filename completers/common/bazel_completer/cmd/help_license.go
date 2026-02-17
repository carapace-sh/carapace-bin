package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_licenseCmd).Standalone()

	helpCmd.AddCommand(help_licenseCmd)
}
