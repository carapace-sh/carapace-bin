package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetCapabilitiesCmd = &cobra.Command{
	Use:   "Get-Capabilities",
	Short: "list capabilities and their install status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetCapabilitiesCmd).Standalone()
	rootCmd.AddCommand(GetCapabilitiesCmd)
}
