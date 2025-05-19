package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flagSCmd = &cobra.Command{
	Use:     "S",
	Aliases: []string{"S"},
	Short:   "Perform a substring search of cask tokens and formula names for <text>",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flagSCmd).Standalone()

	flagSCmd.Flags().Bool("archlinux", false, "Search for <text> in the given database.")
	flagSCmd.Flags().Bool("cask", false, "Search for casks.")
	flagSCmd.Flags().Bool("closed", false, "Search for only closed GitHub pull requests.")
	flagSCmd.Flags().Bool("debian", false, "Search for <text> in the given database.")
	flagSCmd.Flags().Bool("debug", false, "Display any debugging information.")
	flagSCmd.Flags().Bool("desc", false, "Search for formulae with a description matching <text> and casks with a name or description matching <text>.")
	flagSCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to search their descriptions. Implied if `HOMEBREW_EVAL_ALL` is set.")
	flagSCmd.Flags().Bool("fedora", false, "Search for <text> in the given database.")
	flagSCmd.Flags().Bool("fink", false, "Search for <text> in the given database.")
	flagSCmd.Flags().Bool("formula", false, "Search for formulae.")
	flagSCmd.Flags().Bool("help", false, "Show this message.")
	flagSCmd.Flags().Bool("macports", false, "Search for <text> in the given database.")
	flagSCmd.Flags().Bool("open", false, "Search for only open GitHub pull requests.")
	flagSCmd.Flags().Bool("opensuse", false, "Search for <text> in the given database.")
	flagSCmd.Flags().Bool("pull-request", false, "Search for GitHub pull requests containing <text>.")
	flagSCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	flagSCmd.Flags().Bool("repology", false, "Search for <text> in the given database.")
	flagSCmd.Flags().Bool("ubuntu", false, "Search for <text> in the given database.")
	flagSCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
}
