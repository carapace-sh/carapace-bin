package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var authClearResetsCmd = &cobra.Command{
	Use:   "auth:clear-resets [<name>]",
	Short: "Flush expired password reset tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authClearResetsCmd).Standalone()

	rootCmd.AddCommand(authClearResetsCmd)
}
