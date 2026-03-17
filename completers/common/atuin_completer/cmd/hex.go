package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Terminal emulator for atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hexCmd).Standalone()

	hexCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(hexCmd)
}
