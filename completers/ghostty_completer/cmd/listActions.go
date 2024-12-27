package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listActionsCmd = &cobra.Command{
	Use:   "+list-actions",
	Short: "list all the available keybind actions for Ghostty",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listActionsCmd).Standalone()

	listActionsCmd.Flags().Bool("docs", false, "print out the documentation for each action")
	listActionsCmd.Flags().Bool("help", false, "show help")
	rootCmd.AddCommand(listActionsCmd)
}
