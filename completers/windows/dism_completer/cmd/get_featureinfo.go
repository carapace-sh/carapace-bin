package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetFeatureInfoCmd = &cobra.Command{
	Use:   "Get-FeatureInfo",
	Short: "display information about a specific feature",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetFeatureInfoCmd).Standalone()
	rootCmd.AddCommand(GetFeatureInfoCmd)
}
