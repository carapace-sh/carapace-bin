package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:     "prune",
	Short:   "Remove extraneous packages",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()

	pruneCmd.Flags().BoolP("dev", "D", false, "")
	pruneCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pruneCmd.Flags().Bool("ignore-scripts", false, "")
	pruneCmd.Flags().Bool("no-ignore-scripts", false, "Run lifecycle scripts even if scripts are disabled by configuration")
	pruneCmd.Flags().Bool("no-optional", false, "")
	pruneCmd.Flags().Bool("optional", false, "")
	pruneCmd.Flags().BoolP("prod", "P", false, "")
	pruneCmd.Flags().Bool("production", false, "")
	rootCmd.AddCommand(pruneCmd)
}
