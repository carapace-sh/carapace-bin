package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Remove all local metadata keys matching a key or namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clearCmd).Standalone()

	clearCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(clearCmd)
}
