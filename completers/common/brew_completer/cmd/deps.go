package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var depsCmd = &cobra.Command{
	Use:     "deps",
	Short:   "Show dependencies for <formula>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(depsCmd).Standalone()

	depsCmd.Flags().Bool("HEAD", false, "Show dependencies for HEAD version instead of stable version.")
	depsCmd.Flags().Bool("annotate", false, "Mark any build, test, implicit, optional, or recommended dependencies as such in the output.")
	depsCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	depsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	depsCmd.Flags().Bool("direct", false, "Show only the direct dependencies declared in the formula.")
	depsCmd.Flags().Bool("dot", false, "Show text-based graph description in DOT format.")
	depsCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to list their dependencies.")
	depsCmd.Flags().Bool("for-each", false, "Switch into the mode used by the `--eval-all` option, but only list dependencies for each provided <formula>, one formula per line. This is used for debugging the `--installed`/`--eval-all` display mode.")
	depsCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	depsCmd.Flags().Bool("full-name", false, "List dependencies by their full name.")
	depsCmd.Flags().Bool("graph", false, "Show dependencies as a directed graph.")
	depsCmd.Flags().Bool("help", false, "Show this message.")
	depsCmd.Flags().Bool("include-build", false, "Include `:build` dependencies for <formula>.")
	depsCmd.Flags().Bool("include-optional", false, "Include `:optional` dependencies for <formula>.")
	depsCmd.Flags().Bool("include-requirements", false, "Include requirements in addition to dependencies for <formula>.")
	depsCmd.Flags().Bool("include-test", false, "Include `:test` dependencies for <formula> (non-recursive).")
	depsCmd.Flags().Bool("installed", false, "List dependencies for formulae that are currently installed. If <formula> is specified, list only its dependencies that are currently installed.")
	depsCmd.Flags().Bool("missing", false, "Show only missing dependencies.")
	depsCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	depsCmd.Flags().Bool("skip-recommended", false, "Skip `:recommended` dependencies for <formula>.")
	depsCmd.Flags().Bool("topological", false, "Sort dependencies in topological order.")
	depsCmd.Flags().Bool("tree", false, "Show dependencies as a tree. When given multiple formula arguments, show individual trees for each formula.")
	depsCmd.Flags().Bool("union", false, "Show the union of dependencies for multiple <formula>, instead of the intersection.")
	depsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(depsCmd)

	carapace.Gen(depsCmd).PositionalAnyCompletion(
		action.ActionSearch(depsCmd).FilterArgs(),
	)
}
