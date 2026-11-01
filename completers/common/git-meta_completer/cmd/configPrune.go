package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configPruneCmd = &cobra.Command{
	Use:    "config:prune",
	Short:  "Interactively configure auto-prune rules",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configPruneCmd).Standalone()

	configPruneCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(configPruneCmd)
}
