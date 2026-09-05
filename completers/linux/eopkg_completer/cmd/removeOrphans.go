package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeOrphansCmd = &cobra.Command{
	Use:     "remove-orphans",
	Aliases: []string{"rmo"},
	Short:   "remove any packages that were automatically installed and no longer have any dependency relationship",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeOrphansCmd).Standalone()

	removeOrphansCmd.Flags().BoolP("dry-run", "n", false, "only show what would happen, do not actually perform changes")
	removeOrphansCmd.Flags().Bool("ignore-comar", false, "bypass system configuration")
	removeOrphansCmd.Flags().Bool("ignore-safety", false, "ignore safety switch on system.base component")
	removeOrphansCmd.Flags().BoolP("purge", "p", false, "remove files tagged as configuration files too")

	rootCmd.AddCommand(removeOrphansCmd)
}
