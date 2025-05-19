package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var usesCmd = &cobra.Command{
	Use:     "uses",
	Short:   "Show formulae and casks that specify <formula> as a dependency; that is, show dependents of <formula>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(usesCmd).Standalone()

	usesCmd.Flags().Bool("cask", false, "Include only casks.")
	usesCmd.Flags().Bool("debug", false, "Display any debugging information.")
	usesCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to show their dependents.")
	usesCmd.Flags().Bool("formula", false, "Include only formulae.")
	usesCmd.Flags().Bool("help", false, "Show this message.")
	usesCmd.Flags().Bool("include-build", false, "Include formulae that specify <formula> as a `:build` dependency.")
	usesCmd.Flags().Bool("include-optional", false, "Include formulae that specify <formula> as an `:optional` dependency.")
	usesCmd.Flags().Bool("include-test", false, "Include formulae that specify <formula> as a `:test` dependency.")
	usesCmd.Flags().Bool("installed", false, "Only list formulae and casks that are currently installed.")
	usesCmd.Flags().Bool("missing", false, "Only list formulae and casks that are not currently installed.")
	usesCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	usesCmd.Flags().Bool("recursive", false, "Resolve more than one level of dependencies.")
	usesCmd.Flags().Bool("skip-recommended", false, "Skip all formulae that specify <formula> as a `:recommended` dependency.")
	usesCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(usesCmd)
}
