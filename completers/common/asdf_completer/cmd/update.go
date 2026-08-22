package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update asdf to the latest stable release",
	Long:  "Upgrading asdf via asdf update is no longer supported. Please use your OS package manager to upgrade asdf.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	rootCmd.AddCommand(updateCmd)
}
