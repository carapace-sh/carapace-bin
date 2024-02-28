package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var script_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run script in stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(script_runCmd).Standalone()

	script_runCmd.Flags().BoolS("X", "X", false, "Disable all safeguards")
	script_runCmd.Flags().Bool("disable-check-gen-code", false, "Disable outdated generated code check")
	script_runCmd.Flags().Bool("disable-check-git-remote", false, "Disable checking if local default branch is updated with remote")
	script_runCmd.Flags().String("disable-safeguards", "", "Disable safeguards: {all,none,git,git-untracked,git-uncommitted,outdated-code,git-out-of-sync}")
	script_runCmd.Flags().Bool("dry-run", false, "Plan the execution but do not execute it")
	script_runCmd.Flags().Bool("no-recursive", false, "Do not recurse into child stacks")
	scriptCmd.AddCommand(script_runCmd)
}
