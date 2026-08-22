package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:     "watch",
	Short:   "Watch a session (read-only)",
	Aliases: []string{"w"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(watchCmd).Standalone()

	watchCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(watchCmd)
}
