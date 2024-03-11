package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var secret_rmCmd = &cobra.Command{
	Use:     "rm SECRET [SECRET...]",
	Short:   "Remove one or more secrets",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_rmCmd).Standalone()

	secretCmd.AddCommand(secret_rmCmd)

	carapace.Gen(secret_rmCmd).PositionalAnyCompletion(docker.ActionSecrets())
}
