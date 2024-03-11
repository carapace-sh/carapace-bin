package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nameCmd = &cobra.Command{
	Use:   "name [username]",
	Short: "Username stuff",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nameCmd).Standalone()

	rootCmd.AddCommand(nameCmd)
}
