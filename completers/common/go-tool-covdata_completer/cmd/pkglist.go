package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var pkglistCmd = &cobra.Command{
	Use:   "pkglist",
	Short: "output list of package import paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkglistCmd).Standalone()
	pkglistCmd.Flags().SetInterspersed(false)

	pkglistCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "Write CPU profile to specified file")
	pkglistCmd.Flags().BoolS("hw", "hw", false, "Panic on warnings (for stack trace)")
	pkglistCmd.Flags().StringS("i", "i", "", "Input dirs to examine (comma separated)")
	pkglistCmd.Flags().StringS("memprofile", "memprofile", "", "Write memory profile to specified file")
	pkglistCmd.Flags().StringS("memprofilerate", "memprofilerate", "", "Set memprofile sampling rate to value")
	pkglistCmd.Flags().StringS("pkg", "pkg", "", "Restrict output to package(s) matching specified package pattern.")
	pkglistCmd.Flags().StringS("v", "v", "", "Verbose trace output level")
	rootCmd.AddCommand(pkglistCmd)

	carapace.Gen(pkglistCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"i":          carapace.ActionFiles().List(","),
		"memprofile": carapace.ActionFiles(),
		"pkg":        golang.ActionPackages(),
	})
}
