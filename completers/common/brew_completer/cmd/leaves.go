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

	leavesCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	leavesCmd.Flags().BoolP("help", "h", false, "Show this message.")
	leavesCmd.Flags().BoolP("installed-as-dependency", "p", false, "Only list leaves that were installed as dependencies.")
	leavesCmd.Flags().BoolP("installed-on-request", "r", false, "Only list leaves that were manually installed.")
	leavesCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	leavesCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(leavesCmd)
}
