package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "connect to or disconnect from shared resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(useCmd).Standalone()
	rootCmd.AddCommand(useCmd)
}
