package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Read a file containing the kubeadm configuration API and report any validation problems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_validateCmd).Standalone()

	config_validateCmd.Flags().Bool("allow-experimental-api", false, "Allow validation of experimental, unreleased APIs.")
	config_validateCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	configCmd.AddCommand(config_validateCmd)

	carapace.Gen(config_validateCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
