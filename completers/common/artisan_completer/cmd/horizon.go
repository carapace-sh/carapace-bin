package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonCmd = &cobra.Command{
	Use:   "horizon [--environment [ENVIRONMENT]]",
	Short: "Start a master supervisor in the foreground",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonCmd).Standalone()

	horizonCmd.Flags().String("environment", "", "The environment name")
	rootCmd.AddCommand(horizonCmd)
}
