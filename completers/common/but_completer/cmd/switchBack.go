package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var switchBackCmd = &cobra.Command{
	Use:   "switch-back",
	Short: "Switch back to the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchBackCmd).Standalone()

	switchBackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(switchBackCmd)
}
