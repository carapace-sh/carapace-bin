package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var percentCmd = &cobra.Command{
	Use:   "percent",
	Short: "output total percentage of statements covered",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(percentCmd).Standalone()
	percentCmd.Flags().SetInterspersed(false)

	percentCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "Write CPU profile to specified file")
	percentCmd.Flags().BoolS("hw", "hw", false, "Panic on warnings (for stack trace)")
	percentCmd.Flags().StringS("i", "i", "", "Input dirs to examine (comma separated)")
	percentCmd.Flags().StringS("memprofile", "memprofile", "", "Write memory profile to specified file")
	percentCmd.Flags().StringS("memprofilerate", "memprofilerate", "", "Set memprofile sampling rate to value")
	percentCmd.Flags().StringS("o", "o", "", "Output text format to file")
	percentCmd.Flags().StringS("pkg", "pkg", "", "Restrict output to package(s) matching specified package pattern")
	percentCmd.Flags().StringS("v", "v", "", "Verbose trace output level")
	rootCmd.AddCommand(percentCmd)

	carapace.Gen(percentCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"i":          carapace.ActionFiles().List(","),
		"memprofile": carapace.ActionFiles(),
		"o":          carapace.ActionFiles(),
		"pkg":        golang.ActionPackages(),
	})
}
