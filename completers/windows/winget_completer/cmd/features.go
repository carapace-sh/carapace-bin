package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var featuresCmd = &cobra.Command{
	Use:   "features",
	Short: "Shows the status of experimental features",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(featuresCmd).Standalone()

	rootCmd.AddCommand(featuresCmd)
}
