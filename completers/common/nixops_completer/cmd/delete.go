package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()

	deleteCmd.Flags().Bool("all", false, "Delete all deployments")
	deleteCmd.Flags().Bool("force", false, "Force deletion even if resources still exist")

	rootCmd.AddCommand(deleteCmd)
}
