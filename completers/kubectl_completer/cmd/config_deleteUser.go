package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteUserCmd = &cobra.Command{
	Use:   "delete-user NAME",
	Short: "Delete the specified user from the kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteUserCmd).Standalone()

	configCmd.AddCommand(config_deleteUserCmd)
}
