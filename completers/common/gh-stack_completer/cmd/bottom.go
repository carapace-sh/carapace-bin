package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bottomCmd = &cobra.Command{
	Use:   "bottom",
	Short: "Check out the bottom branch of the stack (closest to the trunk)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bottomCmd).Standalone()

	rootCmd.AddCommand(bottomCmd)
}
