package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config [flags]",
	Short:   "Manage glab settings.",
	Aliases: []string{"conf"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().BoolP("global", "g", false, "Use global config file.")
	rootCmd.AddCommand(configCmd)
}
