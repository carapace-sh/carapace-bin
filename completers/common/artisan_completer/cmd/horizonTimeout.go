package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonTimeoutCmd = &cobra.Command{
	Use:   "horizon:timeout [<environment>]",
	Short: "Get the maximum timeout for the given environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonTimeoutCmd).Standalone()

	rootCmd.AddCommand(horizonTimeoutCmd)
}
