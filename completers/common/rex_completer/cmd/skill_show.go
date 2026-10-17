package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Print the guide the skill gives agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillShowCmd).Standalone()

	skillCmd.AddCommand(skillShowCmd)
}
