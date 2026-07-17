package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bsexecCmd = &cobra.Command{
	Use:   "bsexec",
	Short: "Execute a program in another process bootstrap context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bsexecCmd).Standalone()
	rootCmd.AddCommand(bsexecCmd)
	carapace.Gen(bsexecCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
