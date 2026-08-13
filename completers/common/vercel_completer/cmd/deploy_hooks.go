package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deploy_hooksCmd = &cobra.Command{
	Use:     "deploy-hooks",
	Aliases: []string{"deploy-hook"},
	Short:   "Manage deploy hooks for Git-triggered builds",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deploy_hooksCmd).Standalone()

	rootCmd.AddCommand(deploy_hooksCmd)
}
