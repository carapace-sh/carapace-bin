package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var featureCmd = &cobra.Command{
	Use:   "feature",
	Short: "manage features",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(featureCmd).Standalone()
	rootCmd.AddCommand(featureCmd)
}
