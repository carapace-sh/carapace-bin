package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Login, logout, and refresh your authentication",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authCmd).Standalone()
	rootCmd.AddCommand(authCmd)
}
