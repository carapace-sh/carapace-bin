package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var descCmd = &cobra.Command{
	Use:     "desc",
	Short:   "Display <formula>'s name and one-line description",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(descCmd).Standalone()

	descCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	descCmd.Flags().Bool("debug", false, "Display any debugging information.")
	descCmd.Flags().Bool("description", false, "Search just descriptions for <text>. If <text> is flanked by slashes, it is interpreted as a regular expression.")
	descCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to search their descriptions. Implied if `HOMEBREW_EVAL_ALL` is set.")
	descCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	descCmd.Flags().Bool("help", false, "Show this message.")
	descCmd.Flags().Bool("name", false, "Search just names for <text>. If <text> is flanked by slashes, it is interpreted as a regular expression.")
	descCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	descCmd.Flags().Bool("search", false, "Search both names and descriptions for <text>. If <text> is flanked by slashes, it is interpreted as a regular expression.")
	descCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(descCmd)

	carapace.Gen(descCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if descCmd.Flag("eval-all").Changed {
				return action.ActionSearch(descCmd)
			}
			return action.ActionList(uninstallCmd)
		}).FilterArgs(),
	)
}
