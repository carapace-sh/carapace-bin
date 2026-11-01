package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revsCmd = &cobra.Command{
	Use:   "revs",
	Short: "List or count commits matching revision ranges",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revsCmd).Standalone()

	helpCmd.AddCommand(help_revsCmd)
}
