package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Tailscale to the latest/different version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("dry-run", false, "print what update would do without doing it")
	updateCmd.Flags().Bool("yes", false, "update without interactive prompts")
	rootCmd.AddCommand(updateCmd)
}
