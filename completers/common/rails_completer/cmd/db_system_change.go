package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rails"
	"github.com/spf13/cobra"
)

var dbSystemChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change the database adapter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbSystemChangeCmd).Standalone()
	dbSystemChangeCmd.Flags().BoolP("help", "h", false, "Show help")
	dbSystemChangeCmd.Flags().String("to", "", "Target database adapter")

	carapace.Gen(dbSystemChangeCmd).FlagCompletion(carapace.ActionMap{
		"to": rails.ActionDatabaseAdapters(),
	})
}
