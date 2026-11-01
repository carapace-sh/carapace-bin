package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var targetCmd = &cobra.Command{
	Use:     "target",
	Aliases: []string{"targets"},
	Short:   "Manage your Vercel Project's targets (custom environments)",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(targetCmd).Standalone()

	rootCmd.AddCommand(targetCmd)
}
