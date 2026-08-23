package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dumpCmd)
}

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dump the CloudDocs database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpCmd).Standalone()

	dumpCmd.Flags().String("database-path", "", "use the database at <db-path>")
	dumpCmd.Flags().String("output", "", "redirect output to <file-path>")

	carapace.Gen(dumpCmd).FlagCompletion(carapace.ActionMap{
		"database-path": carapace.ActionFiles(),
		"output":        carapace.ActionFiles(),
	})
}