package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livecheckCmd = &cobra.Command{
	Use:     "livecheck",
	Short:   "Check for newer versions of formulae and/or casks from upstream",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livecheckCmd).Standalone()

	livecheckCmd.Flags().Bool("cask", false, "Only check casks.")
	livecheckCmd.Flags().Bool("debug", false, "Display any debugging information.")
	livecheckCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to check them.")
	livecheckCmd.Flags().Bool("formula", false, "Only check formulae.")
	livecheckCmd.Flags().Bool("full-name", false, "Print formulae and casks with fully-qualified names.")
	livecheckCmd.Flags().Bool("help", false, "Show this message.")
	livecheckCmd.Flags().Bool("installed", false, "Check formulae and casks that are currently installed.")
	livecheckCmd.Flags().Bool("json", false, "Output information in JSON format.")
	livecheckCmd.Flags().Bool("newer-only", false, "Show the latest version only if it's newer than the formula/cask.")
	livecheckCmd.Flags().Bool("quiet", false, "Suppress warnings, don't print a progress bar for JSON output.")
	livecheckCmd.Flags().Bool("resources", false, "Also check resources for formulae.")
	livecheckCmd.Flags().Bool("tap", false, "Check formulae and casks within the given tap, specified as <user>`/`<repo>.")
	livecheckCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(livecheckCmd)
}
