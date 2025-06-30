package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bridgeCmd = &cobra.Command{
	Use:   "bridge CMD [OPTIONS]",
	Short: "Convert compose files into another model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bridgeCmd).Standalone()

	rootCmd.AddCommand(bridgeCmd)
}
