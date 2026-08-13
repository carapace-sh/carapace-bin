package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the current project in the Vercel Dashboard",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openCmd).Standalone()

	openCmd.Flags().Bool("yes", false, "Skip confirmation")

	rootCmd.AddCommand(openCmd)
}
