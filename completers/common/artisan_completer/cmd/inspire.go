package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspireCmd = &cobra.Command{
	Use:   "inspire",
	Short: "Display an inspiring quote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspireCmd).Standalone()

	rootCmd.AddCommand(inspireCmd)
}
