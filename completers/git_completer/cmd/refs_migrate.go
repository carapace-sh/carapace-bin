package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var refs_migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate ref store between different formats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refs_migrateCmd).Standalone()

	refs_migrateCmd.Flags().Bool("dry-run", false, "perform the migration, but do not modify the repository")
	refs_migrateCmd.Flags().String("ref-format", "", "the ref format to migrate the ref store to")
	refsCmd.AddCommand(refs_migrateCmd)

	carapace.Gen(refs_migrateCmd).FlagCompletion(carapace.ActionMap{
		"ref-format": carapace.ActionValuesDescribed(
			"files", "loose files with packed-refs",
			"reftable", "reftable format",
		),
	})
}
