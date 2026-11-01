package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetCapabilityInfoCmd = &cobra.Command{
	Use:   "Get-CapabilityInfo",
	Short: "display information about a specific capability",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetCapabilityInfoCmd).Standalone()
	rootCmd.AddCommand(GetCapabilityInfoCmd)
}
