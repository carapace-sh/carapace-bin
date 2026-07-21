package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:     "prune",
	Short:   "Removes extraneous packages",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()

	pruneCmd.Flags().Bool("no-optional", false, "Remove the packages specified in `optionalDependencies`")
	pruneCmd.Flags().Bool("prod", false, "Remove the packages specified in `devDependencies`")
	rootCmd.AddCommand(pruneCmd)
}
