package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var leavesCmd = &cobra.Command{
	Use:     "leaves",
	Short:   "List installed formulae that are not dependencies of another installed formula or cask",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(leavesCmd).Standalone()

	leavesCmd.Flags().Bool("debug", false, "Display any debugging information.")
	leavesCmd.Flags().Bool("help", false, "Show this message.")
	leavesCmd.Flags().Bool("installed-as-dependency", false, "Only list leaves that were installed as dependencies.")
	leavesCmd.Flags().Bool("installed-on-request", false, "Only list leaves that were manually installed.")
	leavesCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	leavesCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(leavesCmd)
}
