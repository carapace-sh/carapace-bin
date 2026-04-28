package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/go_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fix",
	Long:  "https://pkg.go.dev/cmd/fix",
	Short: "Fix finds Go programs that use old APIs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("V", "V", false, "print version and exit")
	rootCmd.Flags().BoolS("any", "any", false, "enable any analysis")
	rootCmd.Flags().BoolS("buildtag", "buildtag", false, "enable buildtag analysis")
	rootCmd.Flags().StringS("c", "c", "", "display offending line with this many lines of context")
	rootCmd.Flags().BoolS("diff", "diff", false, "display diffs instead of rewriting files")
	rootCmd.Flags().BoolS("fix", "fix", false, "apply all suggested fixes")
	rootCmd.Flags().StringS("fixtool", "fixtool", "", "select a different analysis tool")
	rootCmd.Flags().BoolS("flags", "flags", false, "print analyzer flags in JSON")
	rootCmd.Flags().BoolS("fmtappendf", "fmtappendf", false, "enable fmtappendf analysis")
	rootCmd.Flags().StringS("force", "force", "", "force these fixes to run even if the code looks updated")
	rootCmd.Flags().BoolS("forvar", "forvar", false, "enable forvar analysis")
	rootCmd.Flags().StringS("go", "go", "", "go language version for files")
	rootCmd.Flags().BoolS("hostport", "hostport", false, "enable hostport analysis")
	rootCmd.Flags().BoolS("inline", "inline", false, "enable inline analysis")
	rootCmd.Flags().BoolS("json", "json", false, "emit JSON output")
	rootCmd.Flags().BoolS("mapsloop", "mapsloop", false, "enable mapsloop analysis")
	rootCmd.Flags().BoolS("minmax", "minmax", false, "enable minmax analysis")
	rootCmd.Flags().BoolS("newexpr", "newexpr", false, "enable newexpr analysis")
	rootCmd.Flags().BoolS("omitzero", "omitzero", false, "enable omitzero analysis")
	rootCmd.Flags().BoolS("plusbuild", "plusbuild", false, "enable plusbuild analysis")
	rootCmd.Flags().StringS("r", "r", "", "restrict the rewrites to this comma-separated list")
	rootCmd.Flags().BoolS("rangeint", "rangeint", false, "enable rangeint analysis")
	rootCmd.Flags().BoolS("reflecttypefor", "reflecttypefor", false, "enable reflecttypefor analysis")
	rootCmd.Flags().BoolS("slicescontains", "slicescontains", false, "enable slicescontains analysis")
	rootCmd.Flags().BoolS("slicessort", "slicessort", false, "enable slicessort analysis")
	rootCmd.Flags().BoolS("stditerators", "stditerators", false, "enable stditerators analysis")
	rootCmd.Flags().BoolS("stringsbuilder", "stringsbuilder", false, "enable stringsbuilder analysis")
	rootCmd.Flags().BoolS("stringscut", "stringscut", false, "enable stringscut analysis")
	rootCmd.Flags().BoolS("stringscutprefix", "stringscutprefix", false, "enable stringscutprefix analysis")
	rootCmd.Flags().BoolS("stringsseq", "stringsseq", false, "enable stringsseq analysis")
	rootCmd.Flags().BoolS("testingcontext", "testingcontext", false, "enable testingcontext analysis")
	rootCmd.Flags().BoolS("waitgroup", "waitgroup", false, "enable waitgroup analysis")
	common.AddPackageBuildFlags(rootCmd)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"fixtool": carapace.ActionFiles(),
		"force":   golang.ActionRewrites().UniqueList(","),
		"go":      golang.ActionVersions(),
		"r":       golang.ActionRewrites().UniqueList(","),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
