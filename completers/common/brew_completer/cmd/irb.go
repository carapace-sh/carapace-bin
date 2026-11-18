package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var irbCmd = &cobra.Command{
	Use:     "irb",
	Short:   "Enter the interactive Homebrew Ruby shell",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(irbCmd).Standalone()

	irbCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	irbCmd.Flags().Bool("examples", false, "Show several examples.")
	irbCmd.Flags().BoolP("help", "h", false, "Show this message.")
	irbCmd.Flags().Bool("pry", false, "Use Pry instead of IRB. Implied if `HOMEBREW_PRY` is set.")
	irbCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	irbCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(irbCmd)
}
