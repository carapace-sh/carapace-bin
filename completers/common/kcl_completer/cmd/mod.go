package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var modCmd = &cobra.Command{
	Use:   "mod",
	Short: "KCL module management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(modCmd).Standalone()

	modCmd.PersistentFlags().BoolP("quiet", "q", false, "Set the quiet mode (no output)")
	rootCmd.AddCommand(modCmd)
}
