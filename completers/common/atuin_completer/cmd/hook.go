package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Manage AI-agent shell hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hookCmd).Standalone()

	hookCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(hookCmd)
}
