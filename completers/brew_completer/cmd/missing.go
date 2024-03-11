package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var missingCmd = &cobra.Command{
	Use:     "missing",
	Short:   "Check the given <formula> kegs for missing dependencies",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(missingCmd).Standalone()

	missingCmd.Flags().Bool("debug", false, "Display any debugging information.")
	missingCmd.Flags().Bool("help", false, "Show this message.")
	missingCmd.Flags().Bool("hide", false, "Act as if none of the specified <hidden> are installed. <hidden> should be a comma-separated list of formulae.")
	missingCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	missingCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(missingCmd)
}
