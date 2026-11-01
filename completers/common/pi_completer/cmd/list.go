package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed extensions from settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()
	listCmd.Flags().BoolP("approve", "a", false, "Trust project-local files for this command")
	listCmd.Flags().BoolP("no-approve", "na", false, "Ignore project-local files for this command")
	rootCmd.AddCommand(listCmd)
}
