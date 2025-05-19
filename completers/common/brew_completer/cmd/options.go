package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var optionsCmd = &cobra.Command{
	Use:     "options",
	Short:   "Show install options specific to <formula>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(optionsCmd).Standalone()

	optionsCmd.Flags().Bool("command", false, "Show options for the specified <command>.")
	optionsCmd.Flags().Bool("compact", false, "Show all options on a single line separated by spaces.")
	optionsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	optionsCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to show their options.")
	optionsCmd.Flags().Bool("help", false, "Show this message.")
	optionsCmd.Flags().Bool("installed", false, "Show options for formulae that are currently installed.")
	optionsCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	optionsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(optionsCmd)
}
