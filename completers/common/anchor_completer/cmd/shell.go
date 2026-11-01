package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Starts a node shell with an Anchor client setup according to the local config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCmd).Standalone()

	shellCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(shellCmd)
}
