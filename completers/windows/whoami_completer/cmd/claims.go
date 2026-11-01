package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var claimsCmd = &cobra.Command{
	Use:   "claims",
	Short: "display claims for the current user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(claimsCmd).Standalone()
	rootCmd.AddCommand(claimsCmd)
}
