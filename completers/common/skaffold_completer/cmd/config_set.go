package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value in the global Skaffold config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().StringP("config", "c", "", "Path to Skaffold config")
	config_setCmd.Flags().BoolP("global", "g", false, "Set value for global config")
	config_setCmd.Flags().StringP("id", "i", "", "Set value for given survey config")
	config_setCmd.Flags().StringP("kube-context", "k", "", "Kubectl context to set values against")
	config_setCmd.Flags().BoolP("survey", "s", false, "Set value for skaffold survey config")
	config_setCmd.Flag("id").Hidden = true
	config_setCmd.Flag("survey").Hidden = true
	applyCommonFlagCompletion(config_setCmd)
	configCmd.AddCommand(config_setCmd)
}
