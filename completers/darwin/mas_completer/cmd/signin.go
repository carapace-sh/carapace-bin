package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signinCmd = &cobra.Command{
	Use:   "signin",
	Short: "Sign in to Mac App Store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signinCmd).Standalone()
	signinCmd.Flags().String("dialog", "", "Provide account credentials via dialog")
	rootCmd.AddCommand(signinCmd)
}
