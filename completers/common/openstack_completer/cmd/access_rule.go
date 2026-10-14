package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var access_ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_ruleCmd).Standalone()

	accessCmd.AddCommand(access_ruleCmd)
}
