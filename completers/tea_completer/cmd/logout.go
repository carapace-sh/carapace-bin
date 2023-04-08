package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:     "logout",
	Short:   "Log out from a Gitea server",
	GroupID: "SETUP",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()

	rootCmd.AddCommand(logoutCmd)

	// TODO completion
}
