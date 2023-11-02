package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_diffeditCmd = &cobra.Command{
	Use:   "diffedit",
	Short: "Touch up the content changes in a revision with a diff editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_diffeditCmd).Standalone()

	helpCmd.AddCommand(help_diffeditCmd)
}
