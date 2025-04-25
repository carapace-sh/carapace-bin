package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_print_joinDefaultsCmd = &cobra.Command{
	Use:   "join-defaults",
	Short: "Print default join configuration, that can be used for 'kubeadm join'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_print_joinDefaultsCmd).Standalone()

	config_printCmd.AddCommand(config_print_joinDefaultsCmd)
}
