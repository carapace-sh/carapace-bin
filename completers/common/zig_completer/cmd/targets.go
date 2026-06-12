package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var targetsCmd = &cobra.Command{
	Use:   "targets",
	Short: "List available compilation targets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(targetsCmd).Standalone()

	rootCmd.AddCommand(targetsCmd)

	targetsCmd.Flags().BoolP("help", "h", false, "Print help")
}
