package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show the current user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_whoamiCmd).Standalone()

	auth_whoamiCmd.Flags().Bool("show-tokens", false, "Show the access, id, and refresh tokens")
	authCmd.AddCommand(auth_whoamiCmd)
}
