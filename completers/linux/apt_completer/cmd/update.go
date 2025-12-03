package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update list of available packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("allow-insecure-repositories", false, "allow insecure repositories")
	updateCmd.Flags().Bool("list-cleanup", false, "automatically manage the contents of /var/lib/apt/lists")
	updateCmd.Flags().Bool("print-uris", false, "print file URIs")
	rootCmd.AddCommand(updateCmd)
}
