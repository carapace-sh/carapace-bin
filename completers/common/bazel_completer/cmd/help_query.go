package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_queryCmd).Standalone()

	helpCmd.AddCommand(help_queryCmd)
}
