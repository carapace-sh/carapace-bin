package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetFeaturesCmd = &cobra.Command{
	Use:   "Get-Features",
	Short: "display all features in a package or image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetFeaturesCmd).Standalone()
	rootCmd.AddCommand(GetFeaturesCmd)
}
