package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/terramate"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var script_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run script in stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(script_runCmd).Standalone()
	script_runCmd.Flags().SetInterspersed(false)

	script_runCmd.Flags().BoolS("X", "X", false, "Disable all safeguards")
	script_runCmd.Flags().Bool("disable-check-gen-code", false, "Disable outdated generated code check")
	script_runCmd.Flags().Bool("disable-check-git-remote", false, "Disable checking if local default branch is updated with remote")
	script_runCmd.Flags().String("disable-safeguards", "", "Disable safeguards: {all,none,git,git-untracked,git-uncommitted,outdated-code,git-out-of-sync}")
	script_runCmd.Flags().Bool("dry-run", false, "Plan the execution but do not execute it")
	script_runCmd.Flags().Bool("no-recursive", false, "Do not recurse into child stacks")
	scriptCmd.AddCommand(script_runCmd)

	carapace.Gen(script_runCmd).FlagCompletion(carapace.ActionMap{
		"disable-safeguards": terramate.ActionSafeguards().UniqueList(","),
	})

	carapace.Gen(script_runCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)

	carapace.Gen(script_runCmd).DashAnyCompletion(
		carapace.ActionPositional(script_runCmd),
	)
}
