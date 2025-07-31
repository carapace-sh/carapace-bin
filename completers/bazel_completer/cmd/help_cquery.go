package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cqueryCmd = &cobra.Command{
	Use:   "cquery",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cqueryCmd).Standalone()

	helpCmd.AddCommand(help_cqueryCmd)
}
