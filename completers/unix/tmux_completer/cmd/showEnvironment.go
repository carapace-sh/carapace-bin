package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var showEnvironmentCmd = &cobra.Command{
	Use:     "show-environment",
	Aliases: []string{"showenv"},
	Short:   "display the environment",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showEnvironmentCmd).Standalone()

	showEnvironmentCmd.Flags().BoolS("g", "g", false, "show global environment")
	showEnvironmentCmd.Flags().BoolS("h", "h", false, "show hidden variables")
	showEnvironmentCmd.Flags().BoolS("s", "s", false, "format output as Bourne shell commands")
	showEnvironmentCmd.Flags().StringS("t", "t", "", "specify target session")
	rootCmd.AddCommand(showEnvironmentCmd)

	carapace.Gen(showEnvironmentCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionSessions(),
	})
}
