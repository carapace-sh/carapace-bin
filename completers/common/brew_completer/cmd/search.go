package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Perform a substring search of cask tokens and formula names for <text>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().Bool("archlinux", false, "Search for <text> in the given database.")
	searchCmd.Flags().Bool("cask", false, "Search for casks.")
	searchCmd.Flags().Bool("closed", false, "Search for only closed GitHub pull requests.")
	searchCmd.Flags().Bool("debian", false, "Search for <text> in the given database.")
	searchCmd.Flags().Bool("debug", false, "Display any debugging information.")
	searchCmd.Flags().Bool("desc", false, "Search for formulae with a description matching <text> and casks with a name or description matching <text>.")
	searchCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to search their descriptions. Implied if `HOMEBREW_EVAL_ALL` is set.")
	searchCmd.Flags().Bool("fedora", false, "Search for <text> in the given database.")
	searchCmd.Flags().Bool("fink", false, "Search for <text> in the given database.")
	searchCmd.Flags().Bool("formula", false, "Search for formulae.")
	searchCmd.Flags().Bool("help", false, "Show this message.")
	searchCmd.Flags().Bool("macports", false, "Search for <text> in the given database.")
	searchCmd.Flags().Bool("open", false, "Search for only open GitHub pull requests.")
	searchCmd.Flags().Bool("opensuse", false, "Search for <text> in the given database.")
	searchCmd.Flags().Bool("pull-request", false, "Search for GitHub pull requests containing <text>.")
	searchCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	searchCmd.Flags().Bool("repology", false, "Search for <text> in the given database.")
	searchCmd.Flags().Bool("ubuntu", false, "Search for <text> in the given database.")
	searchCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(searchCmd)
}
