package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all stored authentication data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_clearCmd).Standalone()

	auth_clearCmd.Flags().BoolP("help", "h", false, "Print help")
	authCmd.AddCommand(auth_clearCmd)
}
