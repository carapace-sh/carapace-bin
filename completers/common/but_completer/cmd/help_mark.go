package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_markCmd = &cobra.Command{
	Use:   "mark",
	Short: "Creates or removes a rule for auto-assigning or auto-comitting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_markCmd).Standalone()

	helpCmd.AddCommand(help_markCmd)
}
