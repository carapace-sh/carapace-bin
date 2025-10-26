package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_configureDockerCmd = &cobra.Command{
	Use:   "configure-docker",
	Short: "Register glab as a Docker credential helper",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_configureDockerCmd).Standalone()

	authCmd.AddCommand(auth_configureDockerCmd)
}
