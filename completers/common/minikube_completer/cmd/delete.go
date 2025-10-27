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
	deleteCmd.Flags().StringP("output", "o", "", "Format to print stdout in. Options include: [text,json]")
	deleteCmd.Flags().Bool("purge", false, "Set this flag to delete the '.minikube' folder from your user directory.")
	rootCmd.AddCommand(deleteCmd)

	carapace.Gen(deleteCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})
}
