package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getUsersCmd = &cobra.Command{
	Use:   "get-users",
	Short: "Display users defined in the kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getUsersCmd).Standalone()

	configCmd.AddCommand(config_getUsersCmd)
}
