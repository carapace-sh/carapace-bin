package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var teamsCmd = &cobra.Command{
	Use:     "teams",
	Aliases: []string{"team"},
	Short:   "Manage Teams under your Vercel account",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teamsCmd).Standalone()

	rootCmd.AddCommand(teamsCmd)
}
