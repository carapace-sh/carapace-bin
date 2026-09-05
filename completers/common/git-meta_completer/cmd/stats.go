package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show metadata statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statsCmd).Standalone()

	statsCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(statsCmd)
}
