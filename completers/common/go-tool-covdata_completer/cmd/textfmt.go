package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var textfmtCmd = &cobra.Command{
	Use:   "textfmt",
	Short: "convert coverage data to textual format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textfmtCmd).Standalone()
	textfmtCmd.Flags().SetInterspersed(false)

	textfmtCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "Write CPU profile to specified file")
	textfmtCmd.Flags().BoolS("hw", "hw", false, "Panic on warnings (for stack trace)")
	textfmtCmd.Flags().StringS("i", "i", "", "Input dirs to examine (comma separated)")
	textfmtCmd.Flags().StringS("memprofile", "memprofile", "", "Write memory profile to specified file")
	textfmtCmd.Flags().StringS("memprofilerate", "memprofilerate", "", "Set memprofile sampling rate to value")
	textfmtCmd.Flags().StringS("o", "o", "", "Output text format to file")
	textfmtCmd.Flags().StringS("pkg", "pkg", "", "Restrict output to package(s) matching specified package pattern")
	textfmtCmd.Flags().StringS("v", "v", "", "Verbose trace output level")
	rootCmd.AddCommand(textfmtCmd)

	carapace.Gen(textfmtCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"i":          carapace.ActionFiles().List(","),
		"memprofile": carapace.ActionFiles(),
		"o":          carapace.ActionFiles(),
		"pkg":        golang.ActionPackages(),
	})
}
