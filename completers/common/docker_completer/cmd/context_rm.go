package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var context_rmCmd = &cobra.Command{
	Use:     "rm CONTEXT [CONTEXT...]",
	Short:   "Remove one or more contexts",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_rmCmd).Standalone()

	context_rmCmd.Flags().BoolP("force", "f", false, "Force the removal of a context in use")
	contextCmd.AddCommand(context_rmCmd)

	carapace.Gen(context_rmCmd).PositionalAnyCompletion(docker.ActionContexts())
}
