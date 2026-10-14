package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tapCmd = &cobra.Command{
	Use:   "tap",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tapCmd).Standalone()

	rootCmd.AddCommand(tapCmd)
}
