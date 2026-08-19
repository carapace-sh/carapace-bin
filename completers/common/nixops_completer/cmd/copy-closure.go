package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var CopyClosureCmd = &cobra.Command{
	Use:   "copy-closure",
	Short: "Copy Closure",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(CopyClosureCmd).Standalone()
	rootCmd.AddCommand(CopyClosureCmd)
}
