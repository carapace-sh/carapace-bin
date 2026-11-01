package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy data from source to destination",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copyCmd).Standalone()
	rootCmd.AddCommand(copyCmd)
}
