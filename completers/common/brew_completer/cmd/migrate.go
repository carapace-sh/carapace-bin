package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
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
	migrateCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	migrateCmd.Flags().BoolP("dry-run", "n", false, "Show what would be migrated, but do not actually migrate anything.")
	migrateCmd.Flags().BoolP("force", "f", false, "Treat installed <formula> and provided <formula> as if they are from the same taps and migrate them anyway.")
	migrateCmd.Flags().Bool("formula", false, "Only migrate formulae.")
	migrateCmd.Flags().BoolP("help", "h", false, "Show this message.")
	migrateCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	migrateCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(migrateCmd)

	carapace.Gen(migrateCmd).PositionalAnyCompletion(
		action.ActionList(migrateCmd).FilterArgs(),
	)
}
