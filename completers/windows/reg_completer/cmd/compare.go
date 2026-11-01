package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "compare registry subkeys or entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compareCmd).Standalone()
	rootCmd.AddCommand(compareCmd)
}
