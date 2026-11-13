package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all values set in the global Skaffold config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listCmd).Standalone()

	config_listCmd.Flags().BoolP("all", "a", false, "Show values for all kubecontexts")
	config_listCmd.Flags().StringP("config", "c", "", "Path to Skaffold config")
	config_listCmd.Flags().StringP("kube-context", "k", "", "Kubectl context to set values against")
	applyCommonFlagCompletion(config_listCmd)
	configCmd.AddCommand(config_listCmd)
}
