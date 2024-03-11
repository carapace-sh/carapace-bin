package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Deletes a local Kubernetes cluster",
	GroupID: "basic",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()
	deleteCmd.Flags().Bool("all", false, "Set flag to delete all profiles")
	deleteCmd.Flags().Bool("purge", false, "Set this flag to delete the '.minikube' folder from your user directory.")
	rootCmd.AddCommand(deleteCmd)
}
