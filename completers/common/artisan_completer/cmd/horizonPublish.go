package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonPublishCmd = &cobra.Command{
	Use:   "horizon:publish",
	Short: "Publish all of the Horizon resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonPublishCmd).Standalone()

	rootCmd.AddCommand(horizonPublishCmd)
}
