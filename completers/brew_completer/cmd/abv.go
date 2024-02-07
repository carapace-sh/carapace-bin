package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var abvCmd = &cobra.Command{
	Use:   "abv",
	Short: "Display brief statistics for your Homebrew installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(abvCmd).Standalone()

	abvCmd.Flags().Bool("analytics", false, "List global Homebrew analytics data or, if specified, installation and build error data for <formula> (provided neither `HOMEBREW_NO_ANALYTICS` nor `HOMEBREW_NO_GITHUB_API` are set).")
	abvCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	abvCmd.Flags().Bool("category", false, "Which type of analytics data to retrieve. The value for <category> must be `install`, `install-on-request` or `build-error`; `cask-install` or `os-version` may be specified if <formula> is not. The default is `install`.")
	abvCmd.Flags().Bool("days", false, "How many days of analytics data to retrieve. The value for <days> must be `30`, `90` or `365`. The default is `30`.")
	abvCmd.Flags().Bool("debug", false, "Display any debugging information.")
	abvCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to print their JSON. Implied if `HOMEBREW_EVAL_ALL` is set.")
	abvCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	abvCmd.Flags().Bool("github", false, "Open the GitHub source page for <formula> and <cask> in a browser. To view the history locally: `brew log -p` <formula> or <cask>")
	abvCmd.Flags().Bool("help", false, "Show this message.")
	abvCmd.Flags().Bool("installed", false, "Print JSON of formulae that are currently installed.")
	abvCmd.Flags().Bool("json", false, "Print a JSON representation. Currently the default value for <version> is `v1` for <formula>. For <formula> and <cask> use `v2`. See the docs for examples of using the JSON output: <https://docs.brew.sh/Querying-Brew>")
	abvCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	abvCmd.Flags().Bool("variations", false, "Include the variations hash in each formula's JSON output.")
	abvCmd.Flags().Bool("verbose", false, "Show more verbose analytics data for <formula>.")
	rootCmd.AddCommand(abvCmd)
}
