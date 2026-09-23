package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Start monitoring for crashes or ANRs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitorCmd).Standalone()

	monitorCmd.Flags().BoolP("continue", "c", false, "assume the input is always continue")
	monitorCmd.Flags().StringP("gdb", "", "", "start gdbserv on the given `PORT` at crash/ANR")
	monitorCmd.Flags().BoolP("kill", "k", false, "assume the input is always kill")
	monitorCmd.Flags().StringP("process", "p", "", "only show events related to a specific process or `PACKAGE`")
	monitorCmd.Flags().BoolP("simple", "s", false, "simple mode, only show a summary line for each event")

	rootCmd.AddCommand(monitorCmd)

	carapace.Gen(monitorCmd).FlagCompletion(carapace.ActionMap{
		"process": android.ActionPackages(),
	})

}
