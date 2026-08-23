package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "perform an update of XProtect assets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().Bool("prerelease", false, "Attempt to use a prerelease update")
}