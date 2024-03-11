package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var funcCmd = &cobra.Command{
	Use:   "func",
	Short: "output coverage profile information for each function",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(funcCmd).Standalone()
	funcCmd.Flags().SetInterspersed(false)

	funcCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "Write CPU profile to specified file")
	funcCmd.Flags().BoolS("hw", "hw", false, "Panic on warnings (for stack trace)")
	funcCmd.Flags().StringS("i", "i", "", "Input dirs to examine (comma separated)")
	funcCmd.Flags().StringS("memprofile", "memprofile", "", "Write memory profile to specified file")
	funcCmd.Flags().StringS("memprofilerate", "memprofilerate", "", "Set memprofile sampling rate to value")
	funcCmd.Flags().StringS("pkg", "pkg", "", "Restrict output to package(s) matching specified package pattern")
	funcCmd.Flags().StringS("v", "v", "", "Verbose trace output level")
	rootCmd.AddCommand(funcCmd)

	carapace.Gen(funcCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"i":          carapace.ActionFiles().List(","),
		"memprofile": carapace.ActionFiles(),
		"pkg":        golang.ActionPackages(),
	})
}
