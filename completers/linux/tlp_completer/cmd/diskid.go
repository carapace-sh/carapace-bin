package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diskidCmd = &cobra.Command{
	Use:   "diskid",
	Short: "Print disk ids for configured drives",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diskidCmd).Standalone()

	rootCmd.AddCommand(diskidCmd)
}
