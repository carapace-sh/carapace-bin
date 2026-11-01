package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var continueCmd = &cobra.Command{
	Use:   "continue",
	Short: "continue a paused service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(continueCmd).Standalone()
	rootCmd.AddCommand(continueCmd)
}
