package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeSessionTableCmd = &cobra.Command{
	Use:     "make:session-table",
	Short:   "Create a migration for the session database table",
	Aliases: []string{"session:table"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeSessionTableCmd).Standalone()

	rootCmd.AddCommand(makeSessionTableCmd)
}
