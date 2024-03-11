package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "merge data files together",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeCmd).Standalone()
	mergeCmd.Flags().SetInterspersed(false)

	mergeCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "Write CPU profile to specified file")
	mergeCmd.Flags().BoolS("hw", "hw", false, "Panic on warnings (for stack trace)")
	mergeCmd.Flags().StringS("i", "i", "", "Input dirs to examine (comma separated)")
	mergeCmd.Flags().StringS("memprofile", "memprofile", "", "Write memory profile to specified file")
	mergeCmd.Flags().StringS("memprofilerate", "memprofilerate", "", "Set memprofile sampling rate to value")
	mergeCmd.Flags().StringS("o", "o", "", "Output directory to write")
	mergeCmd.Flags().BoolS("pcombine", "pcombine", false, "Combine profiles derived from distinct program executables")
	mergeCmd.Flags().StringS("pkg", "pkg", "", "Restrict output to package(s) matching specified package pattern")
	mergeCmd.Flags().StringS("v", "v", "", "Verbose trace output level")
	rootCmd.AddCommand(mergeCmd)

	carapace.Gen(mergeCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"i":          carapace.ActionFiles().List(","),
		"memprofile": carapace.ActionFiles(),
		"o":          carapace.ActionFiles(),
		"pkg":        golang.ActionPackages(),
	})
}
