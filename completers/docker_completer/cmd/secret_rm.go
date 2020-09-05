package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secret_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_rmCmd).Standalone()

	secretCmd.AddCommand(secret_rmCmd)

	carapace.Gen(secret_rmCmd).PositionalAnyCompletion(action.ActionSecrets())
}
