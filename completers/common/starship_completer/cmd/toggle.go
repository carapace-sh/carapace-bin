package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle a given starship module",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toggleCmd).Standalone()

	toggleCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(toggleCmd)
}
