package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var intersectCmd = &cobra.Command{
	Use:   "intersect",
	Short: "generate intersection of two sets of data files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(intersectCmd).Standalone()
	intersectCmd.Flags().SetInterspersed(false)

	intersectCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "Write CPU profile to specified file")
	intersectCmd.Flags().BoolS("hw", "hw", false, "Panic on warnings (for stack trace)")
	intersectCmd.Flags().StringS("i", "i", "", "Input dirs to examine (comma separated)")
	intersectCmd.Flags().StringS("memprofile", "memprofile", "", "Write memory profile to specified file")
	intersectCmd.Flags().StringS("memprofilerate", "memprofilerate", "", "Set memprofile sampling rate to value")
	intersectCmd.Flags().StringS("o", "o", "", "Output directory to write")
	intersectCmd.Flags().StringS("pkg", "pkg", "", "Restrict output to package(s) matching specified package pattern")
	intersectCmd.Flags().StringS("v", "v", "", "Verbose trace output level")
	rootCmd.AddCommand(intersectCmd)

	carapace.Gen(intersectCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"i":          carapace.ActionFiles().List(","),
		"memprofile": carapace.ActionFiles(),
		"o":          carapace.ActionFiles(),
		"pkg":        golang.ActionPackages(),
	})
}
