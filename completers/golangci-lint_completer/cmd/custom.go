package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Build a version of golangci-lint with custom linters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customCmd).Standalone()

	rootCmd.AddCommand(customCmd)
}
