package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shimversionsCmd = &cobra.Command{
	Use:   "shimversions",
	Short: "List the plugins and versions that provide a command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shimversionsCmd).Standalone()

	rootCmd.AddCommand(shimversionsCmd)
}
