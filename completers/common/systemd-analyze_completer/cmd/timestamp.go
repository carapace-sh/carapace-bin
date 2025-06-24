package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestampCmd = &cobra.Command{
	Use:   "timestamp",
	Short: "Validate a timestamp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestampCmd).Standalone()

	rootCmd.AddCommand(timestampCmd)
}
