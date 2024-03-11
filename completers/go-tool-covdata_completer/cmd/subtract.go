package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "subtract one set of data files from another set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subtractCmd).Standalone()
	subtractCmd.Flags().SetInterspersed(false)

	subtractCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "Write CPU profile to specified file")
	subtractCmd.Flags().BoolS("hw", "hw", false, "Panic on warnings (for stack trace)")
	subtractCmd.Flags().StringS("i", "i", "", "Input dirs to examine (comma separated)")
	subtractCmd.Flags().StringS("memprofile", "memprofile", "", "Write memory profile to specified file")
	subtractCmd.Flags().StringS("memprofilerate", "memprofilerate", "", "Set memprofile sampling rate to value")
	subtractCmd.Flags().StringS("o", "o", "", "Output directory to write")
	subtractCmd.Flags().StringS("pkg", "pkg", "", "Restrict output to package(s) matching specified package pattern")
	subtractCmd.Flags().StringS("v", "v", "", "Verbose trace output level")
	rootCmd.AddCommand(subtractCmd)

	carapace.Gen(subtractCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"i":          carapace.ActionFiles().List(","),
		"memprofile": carapace.ActionFiles(),
		"o":          carapace.ActionFiles(),
		"pkg":        golang.ActionPackages(),
	})
}
