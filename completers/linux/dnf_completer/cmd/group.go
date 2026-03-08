package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupCmd = &cobra.Command{
	Use:   "group <subcommand> [options] [<group-spec>...]",
	Short: "manage comps groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupCmd).Standalone()

	rootCmd.AddCommand(groupCmd)
}
