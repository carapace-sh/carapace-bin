package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:     "exec",
	Short:   "Run a command from a local or remote npm package",
	Aliases: []string{"x"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()
	execCmd.Flags().StringArray("allow-scripts", nil, "packages whose install-time lifecycle scripts are allowed to run")
	execCmd.Flags().String("call", "", "optional companion option")
	execCmd.Flags().Bool("dangerously-allow-all-scripts", false, "bypass the allowScripts policy entirely")
	execCmd.Flags().String("package", "", "package to install for exec")
	execCmd.Flags().Bool("strict-allow-scripts", false, "turn install-script policy from warning into error")
	addWorkspaceFlags(execCmd)

	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"package": action.ActionPackages(execCmd),
	})

	carapace.Gen(execCmd).PositionalCompletion(
		action.ActionPackages(execCmd),
	)
}
