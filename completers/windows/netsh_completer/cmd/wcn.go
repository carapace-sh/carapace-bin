package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wcnCmd = &cobra.Command{
	Use:   "wcn",
	Short: "Windows Connect Now configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wcnCmd).Standalone()
	rootCmd.AddCommand(wcnCmd)
}
