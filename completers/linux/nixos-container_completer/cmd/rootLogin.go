package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootLoginCmd = &cobra.Command{
	Use:   "root-login",
	Short: "open an interactive root shell in the container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootLoginCmd).Standalone()
	rootCmd.AddCommand(rootLoginCmd)
}
