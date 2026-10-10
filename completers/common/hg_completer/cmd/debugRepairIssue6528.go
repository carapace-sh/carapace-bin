package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugRepairIssue6528Cmd = &cobra.Command{
	Use:    "debug-repair-issue6528",
	Short:  "find affected revisions and repair them.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugRepairIssue6528Cmd).Standalone()

	debugRepairIssue6528Cmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	debugRepairIssue6528Cmd.Flags().String("from-report", "", "repair revisions listed in this report file")
	debugRepairIssue6528Cmd.Flags().Bool("paranoid", false, "check that both detection methods do the same thing")
	debugRepairIssue6528Cmd.Flags().String("to-report", "", "build a report of affected revisions to this file")
	rootCmd.AddCommand(debugRepairIssue6528Cmd)
}
