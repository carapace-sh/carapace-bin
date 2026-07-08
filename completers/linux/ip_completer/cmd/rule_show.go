package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rule_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list", "lst"},
	Short:   "list rules",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rule_showCmd).Standalone()

	ruleCmd.AddCommand(rule_showCmd)
}
