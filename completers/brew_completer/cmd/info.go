package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Display brief statistics for your Homebrew installation",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().Bool("analytics", false, "List global Homebrew analytics data or, if specified, installation and build error data for <formula> (provided neither `HOMEBREW_NO_ANALYTICS` nor `HOMEBREW_NO_GITHUB_API` are set).")
	infoCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	infoCmd.Flags().String("category", "", "Which type of analytics data to retrieve. The value for <category> must be `install`, `install-on-request` or `build-error`; `cask-install` or `os-version` may be specified if <formula> is not. The default is `install`.")
	infoCmd.Flags().Bool("days", false, "How many days of analytics data to retrieve. The value for <days> must be `30`, `90` or `365`. The default is `30`.")
	infoCmd.Flags().Bool("debug", false, "Display any debugging information.")
	infoCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to print their JSON. Implied if `HOMEBREW_EVAL_ALL` is set.")
	infoCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	infoCmd.Flags().Bool("github", false, "Open the GitHub source page for <formula> and <cask> in a browser. To view the history locally: `brew log -p` <formula> or <cask>")
	infoCmd.Flags().Bool("help", false, "Show this message.")
	infoCmd.Flags().Bool("installed", false, "Print JSON of formulae that are currently installed.")
	infoCmd.Flags().Bool("json", false, "Print a JSON representation. Currently the default value for <version> is `v1` for <formula>. For <formula> and <cask> use `v2`. See the docs for examples of using the JSON output: <https://docs.brew.sh/Querying-Brew>")
	infoCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	infoCmd.Flags().Bool("variations", false, "Include the variations hash in each formula's JSON output.")
	infoCmd.Flags().Bool("verbose", false, "Show more verbose analytics data for <formula>.")
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"category": carapace.ActionValues("install", "install-on-request", "build-error"),
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		action.ActionSearch(installCmd),
	)
}
