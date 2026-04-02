package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeOrphansCmd = &cobra.Command{
	Use:     "remove-orphans",
	Short:   "Remove orphaned packages",
	Aliases: []string{"rmo"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeOrphansCmd).Standalone()

	removeOrphansCmd.Flags().Bool("ignore-dependency", false, "Do not take dependency info into account")
	removeOrphansCmd.Flags().Bool("ignore-comar", false, "Bypass COMAR configuration agent")
	removeOrphansCmd.Flags().Bool("ignore-safety", false, "Bypass safety switch")
	removeOrphansCmd.Flags().BoolP("dry-run", "n", false, "Do not perform any action, just show what would be done")
	removeOrphansCmd.Flags().Bool("purge", false, "Remove package and purge")

	rootCmd.AddCommand(removeOrphansCmd)
}
