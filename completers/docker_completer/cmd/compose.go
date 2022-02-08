package cmd

import (
	"os"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var composeCmd = &cobra.Command{
	Use:                "compose",
	Short:              "Define and run multi-container applications with Docker",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(composeCmd).Standalone()

	rootCmd.AddCommand(composeCmd)

	carapace.Gen(composeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO experimental - this prevents building the completer separately
			// maybe replace with invocation of cmd.Execute() and catch stdout like in the root command of carapace-bin
			executable, err := os.Executable()
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			arg := []string{"docker-compose", "export", "_", "docker-compose"}
			arg = append(arg, c.Args...)
			arg = append(arg, c.CallbackValue)
			return carapace.ActionExecCarapace(executable, arg...)
		}),
	)
}
