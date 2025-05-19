package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Shows the username of the currently logged in user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whoamiCmd).Standalone()

	rootCmd.AddCommand(whoamiCmd)
}
