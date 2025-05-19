package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_print_upgradeDefaultsCmd = &cobra.Command{
	Use:   "upgrade-defaults",
	Short: "Print default upgrade configuration, that can be used for 'kubeadm upgrade'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_print_upgradeDefaultsCmd).Standalone()

	config_printCmd.AddCommand(config_print_upgradeDefaultsCmd)
}
