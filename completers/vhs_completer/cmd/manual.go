package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manualCmd = &cobra.Command{
	Use:     "manual",
	Short:   "Generate man pages",
	Aliases: []string{"man"},
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manualCmd).Standalone()

	rootCmd.AddCommand(manualCmd)
}
