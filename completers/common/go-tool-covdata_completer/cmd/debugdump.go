package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var debugdumpCmd = &cobra.Command{
	Use:   "debugdump",
	Short: "dump data in human-readable format for debugging purposes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugdumpCmd).Standalone()
	debugdumpCmd.Flags().SetInterspersed(false)

	debugdumpCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "Write CPU profile to specified file")
	debugdumpCmd.Flags().BoolS("hw", "hw", false, "Panic on warnings (for stack trace)")
	debugdumpCmd.Flags().StringS("i", "i", "", "Input dirs to examine (comma separated)")
	debugdumpCmd.Flags().BoolS("live", "live", false, "Select only live (executed) functions for dump output")
	debugdumpCmd.Flags().StringS("memprofile", "memprofile", "", "Write memory profile to specified file")
	debugdumpCmd.Flags().StringS("memprofilerate", "memprofilerate", "", "Set memprofile sampling rate to value")
	debugdumpCmd.Flags().StringS("pkg", "pkg", "", "Restrict output to package(s) matching specified package pattern")
	debugdumpCmd.Flags().StringS("v", "v", "", "Verbose trace output level")
	rootCmd.AddCommand(debugdumpCmd)

	carapace.Gen(debugdumpCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"i":          carapace.ActionFiles().List(","),
		"memprofile": carapace.ActionFiles(),
		"pkg":        golang.ActionPackages(),
	})
}
