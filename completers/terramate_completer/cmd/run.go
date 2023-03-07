package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util/embed"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run command in the stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("continue-on-error", false, "Continue executing in other stacks in case of error")
	runCmd.Flags().Bool("disable-check-gen-code", false, "Disable outdated generated code check")
	runCmd.Flags().Bool("disable-check-git-remote", false, "Disable checking if local default branch is updated with remote")
	runCmd.Flags().Bool("dry-run", false, "Plan the execution but do not execute it")
	runCmd.Flags().Bool("no-recursive", false, "Do not recurse into child stacks")
	runCmd.Flags().Bool("reverse", false, "Reverse the order of execution")
	rootCmd.AddCommand(runCmd)

	embed.EmbedCarapaceBin(runCmd)
}
