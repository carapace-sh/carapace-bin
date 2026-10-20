package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setupTokenCmd = &cobra.Command{
	Use:   "setup-token",
	Short: "Set up a long-lived authentication token (requires Claude subscription)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupTokenCmd).Standalone()

	setupTokenCmd.Flags().BoolP("help", "h", false, "Display help for command")
	rootCmd.AddCommand(setupTokenCmd)
}
