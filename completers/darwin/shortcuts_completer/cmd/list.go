package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all shortcuts or shortcut folders",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()
	listCmd.Flags().StringP("folder", "f", "", "List shortcuts only from a specific folder")
	listCmd.Flags().Bool("folders", false, "List all folder names instead of shortcuts")
	rootCmd.AddCommand(listCmd)
}
