package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "exec",
	Short: "Replace the shell with the given command",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-exec",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().StringS("a", "a", "", "pass name as the zeroth argument to command")
	rootCmd.Flags().BoolS("c", "c", false, "execute command with an empty environment")
	rootCmd.Flags().BoolS("l", "l", false, "place a dash in the zeroth argument to command")

	carapace.Gen(rootCmd).PositionalCompletion(
		shell.ActionExecutables(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
