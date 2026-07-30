package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var DisableFeatureCmd = &cobra.Command{
	Use:   "Disable-Feature",
	Short: "disable a feature in an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(DisableFeatureCmd).Standalone()
	rootCmd.AddCommand(DisableFeatureCmd)
}
