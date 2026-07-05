package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signoutCmd = &cobra.Command{
	Use:   "signout",
	Short: "Sign out of Mac App Store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signoutCmd).Standalone()
	rootCmd.AddCommand(signoutCmd)
}
