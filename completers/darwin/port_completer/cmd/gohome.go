package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gohomeCmd = &cobra.Command{
	Use:   "gohome",
	Short: "Open the home page for a port in the default browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gohomeCmd).Standalone()
	rootCmd.AddCommand(gohomeCmd)
}
