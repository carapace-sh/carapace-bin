package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap a service into the specified domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrapCmd).Standalone()
	rootCmd.AddCommand(bootstrapCmd)
	carapace.Gen(bootstrapCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
