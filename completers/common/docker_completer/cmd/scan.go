package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Docker Scan",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scanCmd).Standalone()

	rootCmd.AddCommand(scanCmd)
}
