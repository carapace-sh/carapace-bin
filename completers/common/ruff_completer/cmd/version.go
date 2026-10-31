package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display Ruff's version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	versionCmd.Flags().String("output-format", "text", "")
	rootCmd.AddCommand(versionCmd)
}
