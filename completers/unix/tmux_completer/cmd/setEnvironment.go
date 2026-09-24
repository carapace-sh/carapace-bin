package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var setEnvironmentCmd = &cobra.Command{
	Use:     "set-environment",
	Aliases: []string{"setenv"},
	Short:   "(un)set an environment variable",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setEnvironmentCmd).Standalone()

	setEnvironmentCmd.Flags().BoolS("F", "F", false, "expand value as a format")
	setEnvironmentCmd.Flags().BoolS("g", "g", false, "modify global environment")
	setEnvironmentCmd.Flags().BoolS("h", "h", false, "mark the variable as hidden")
	setEnvironmentCmd.Flags().BoolS("r", "r", false, "remove variable before starting new processes")
	setEnvironmentCmd.Flags().StringS("t", "t", "", "specify target session")
	setEnvironmentCmd.Flags().BoolS("u", "u", false, "unset a variable")
	rootCmd.AddCommand(setEnvironmentCmd)

	carapace.Gen(setEnvironmentCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionSessions(),
	})
}
