package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_print_resetDefaultsCmd = &cobra.Command{
	Use:   "reset-defaults",
	Short: "Print default reset configuration, that can be used for 'kubeadm reset'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_print_resetDefaultsCmd).Standalone()

	config_printCmd.AddCommand(config_print_resetDefaultsCmd)
}
