package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List apps installed from the App Store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()
	listCmd.Flags().Bool("bundle", false, "Process all app IDs as bundle IDs")
	listCmd.Flags().Bool("json", false, "Output JSON")
	rootCmd.AddCommand(listCmd)
}
