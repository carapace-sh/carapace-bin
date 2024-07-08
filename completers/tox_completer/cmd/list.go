package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"l"},
	Short: "list environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	add_common_flags(listCmd)
	add_env_filtering_flags(listCmd)

	listCmd.Flags().BoolS("d", "d", false, "list just default envs (default: False)")

	rootCmd.AddCommand(listCmd)
}
