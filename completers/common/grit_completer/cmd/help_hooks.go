package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "List hook scripts or run one by name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_hooksCmd).Standalone()

	helpCmd.AddCommand(help_hooksCmd)
}
