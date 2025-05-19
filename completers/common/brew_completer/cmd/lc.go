package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lcCmd = &cobra.Command{
	Use:   "lc",
	Short: "Check for newer versions of formulae and/or casks from upstream",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lcCmd).Standalone()

	lcCmd.Flags().Bool("cask", false, "Only check casks.")
	lcCmd.Flags().Bool("debug", false, "Display any debugging information.")
	lcCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to check them.")
	lcCmd.Flags().Bool("formula", false, "Only check formulae.")
	lcCmd.Flags().Bool("full-name", false, "Print formulae and casks with fully-qualified names.")
	lcCmd.Flags().Bool("help", false, "Show this message.")
	lcCmd.Flags().Bool("installed", false, "Check formulae and casks that are currently installed.")
	lcCmd.Flags().Bool("json", false, "Output information in JSON format.")
	lcCmd.Flags().Bool("newer-only", false, "Show the latest version only if it's newer than the formula/cask.")
	lcCmd.Flags().Bool("quiet", false, "Suppress warnings, don't print a progress bar for JSON output.")
	lcCmd.Flags().Bool("resources", false, "Also check resources for formulae.")
	lcCmd.Flags().Bool("tap", false, "Check formulae and casks within the given tap, specified as <user>`/`<repo>.")
	lcCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(lcCmd)
}
