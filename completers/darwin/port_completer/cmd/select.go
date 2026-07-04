package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a default version for a group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectCmd).Standalone()
	selectCmd.Flags().Bool("list", false, "List available versions for a group")
	selectCmd.Flags().Bool("set", false, "Set a version for a group")
	selectCmd.Flags().Bool("show", false, "Show the current version for a group")
	selectCmd.Flags().Bool("summary", false, "Show a summary of all groups")
	rootCmd.AddCommand(selectCmd)
}
