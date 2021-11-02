package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_deleteCmd).Standalone()
	apps_deleteCmd.Flags().BoolP("force", "f", false, "Delete the App without a confirmation prompt")
	appsCmd.AddCommand(apps_deleteCmd)

	// TODO app-id completion
}
