package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bottleCmd = &cobra.Command{
	Use:     "bottle",
	Short:   "Generate a bottle (binary package) from a formula that was installed with `--build-bottle`",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bottleCmd).Standalone()

	bottleCmd.Flags().Bool("committer", false, "Specify a committer name and email in `git`'s standard author format.")
	bottleCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bottleCmd.Flags().Bool("force-core-tap", false, "Build a bottle even if <formula> is not in `homebrew/core` or any installed taps.")
	bottleCmd.Flags().Bool("help", false, "Show this message.")
	bottleCmd.Flags().Bool("json", false, "Write bottle information to a JSON file, which can be used as the value for `--merge`.")
	bottleCmd.Flags().Bool("keep-old", false, "If the formula specifies a rebuild version, attempt to preserve its value in the generated DSL.")
	bottleCmd.Flags().Bool("merge", false, "Generate an updated bottle block for a formula and optionally merge it into the formula file. Instead of a formula name, requires the path to a JSON file generated with `brew bottle --json` <formula>.")
	bottleCmd.Flags().Bool("no-all-checks", false, "Don't try to create an `all` bottle or stop a no-change upload.")
	bottleCmd.Flags().Bool("no-commit", false, "When passed with `--write`, a new commit will not generated after writing changes to the formula file.")
	bottleCmd.Flags().Bool("no-rebuild", false, "If the formula specifies a rebuild version, remove it from the generated DSL.")
	bottleCmd.Flags().Bool("only-json-tab", false, "When passed with `--json`, the tab will be written to the JSON file but not the bottle.")
	bottleCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bottleCmd.Flags().Bool("root-url", false, "Use the specified <URL> as the root of the bottle's URL instead of Homebrew's default.")
	bottleCmd.Flags().Bool("root-url-using", false, "Use the specified download strategy class for downloading the bottle's URL instead of Homebrew's default.")
	bottleCmd.Flags().Bool("skip-relocation", false, "Do not check if the bottle can be marked as relocatable.")
	bottleCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bottleCmd.Flags().Bool("write", false, "Write changes to the formula file. A new commit will be generated unless `--no-commit` is passed.")
	rootCmd.AddCommand(bottleCmd)
}
