package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetPESettingsCmd = &cobra.Command{
	Use:   "Get-PESettings",
	Short: "list all WinPE settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetPESettingsCmd).Standalone()
	rootCmd.AddCommand(GetPESettingsCmd)
}
