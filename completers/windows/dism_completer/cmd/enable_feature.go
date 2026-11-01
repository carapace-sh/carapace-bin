package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var EnableFeatureCmd = &cobra.Command{
	Use:   "Enable-Feature",
	Short: "enable a feature in an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(EnableFeatureCmd).Standalone()
	rootCmd.AddCommand(EnableFeatureCmd)
}
