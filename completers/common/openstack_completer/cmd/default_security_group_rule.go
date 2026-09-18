package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var default_security_group_ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(default_security_group_ruleCmd).Standalone()

	default_security_groupCmd.AddCommand(default_security_group_ruleCmd)
}
