package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_interdiffCmd = &cobra.Command{
	Use:   "interdiff",
	Short: "Compare the changes of two commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_interdiffCmd).Standalone()

	helpCmd.AddCommand(help_interdiffCmd)
}
