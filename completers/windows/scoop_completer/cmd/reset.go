package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset an app to resolve conflicts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetCmd).Standalone()
	resetCmd.Flags().BoolP("all", "a", false, "reset all apps (alternative to `*`)")
	rootCmd.AddCommand(resetCmd)
}
