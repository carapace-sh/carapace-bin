package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the active and installed toolchains or profiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().BoolP("help", "h", false, "Print help")
	showCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output with rustc information for all installed toolchains")
	rootCmd.AddCommand(showCmd)
}
