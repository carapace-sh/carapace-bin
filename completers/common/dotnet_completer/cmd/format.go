package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "format code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatCmd).Standalone()
	rootCmd.AddCommand(formatCmd)
}
