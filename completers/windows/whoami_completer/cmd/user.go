package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "display current user information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userCmd).Standalone()
	rootCmd.AddCommand(userCmd)
}
