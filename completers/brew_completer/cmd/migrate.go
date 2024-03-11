package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "Migrate renamed packages to new names, where <formula> are old names of packages",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateCmd).Standalone()

	migrateCmd.Flags().Bool("cask", false, "Only migrate casks.")
	migrateCmd.Flags().Bool("debug", false, "Display any debugging information.")
	migrateCmd.Flags().Bool("dry-run", false, "Show what would be migrated, but do not actually migrate anything.")
	migrateCmd.Flags().Bool("force", false, "Treat installed <formula> and provided <formula> as if they are from the same taps and migrate them anyway.")
	migrateCmd.Flags().Bool("formula", false, "Only migrate formulae.")
	migrateCmd.Flags().Bool("help", false, "Show this message.")
	migrateCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	migrateCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(migrateCmd)
}
