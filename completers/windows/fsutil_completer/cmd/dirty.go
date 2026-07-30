package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dirtyCmd = &cobra.Command{
	Use:   "dirty",
	Short: "query or set a volume's dirty bit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dirtyCmd).Standalone()
	rootCmd.AddCommand(dirtyCmd)
}
