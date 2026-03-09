package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineCmd = &cobra.Command{
	Use:   "offline <subcommand> [options]",
	Short: "manage offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineCmd).Standalone()

	rootCmd.AddCommand(offlineCmd)
}
