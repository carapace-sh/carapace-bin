package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tox"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:     "exec",
	Aliases: []string{"e"},
	Short:   "execute an arbitrary command within a tox environment",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().StringS("e", "e", "", "environment to run (default: py)")
	execCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: '')")
	addCommonSubcommandFlags(execCmd)
	addPkgOnlyFlags(execCmd)
	addCommonRunFlags(execCmd)
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"e": tox.ActionEnvironments(),
	})

	carapace.Gen(execCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
