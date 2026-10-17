package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Inspect and call methods exposed by the server's control API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiCmd).Standalone()

	rootCmd.AddCommand(apiCmd)
}
