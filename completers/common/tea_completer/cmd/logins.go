package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginsCmd = &cobra.Command{
	Use:     "logins",
	Short:   "Log in to a Gitea server",
	GroupID: "SETUP",
	Aliases: []string{"login"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginsCmd).Standalone()

	rootCmd.AddCommand(loginsCmd)
}
