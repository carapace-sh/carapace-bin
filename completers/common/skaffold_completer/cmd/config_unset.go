package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a value in the global Skaffold config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_unsetCmd).Standalone()

	config_unsetCmd.Flags().StringP("config", "c", "", "Path to Skaffold config")
	config_unsetCmd.Flags().BoolP("global", "g", false, "Set value for global config")
	config_unsetCmd.Flags().StringP("id", "i", "", "Set value for given survey config")
	config_unsetCmd.Flags().StringP("kube-context", "k", "", "Kubectl context to set values against")
	config_unsetCmd.Flags().BoolP("survey", "s", false, "Set value for skaffold survey config")
	config_unsetCmd.Flag("id").Hidden = true
	config_unsetCmd.Flag("survey").Hidden = true
	applyCommonFlagCompletion(config_unsetCmd)
	configCmd.AddCommand(config_unsetCmd)
}
