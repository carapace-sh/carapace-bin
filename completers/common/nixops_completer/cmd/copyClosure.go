package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var copyClosureCmd = &cobra.Command{
	Use:   "copy-closure",
	Short: "copy closure to a target machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copyClosureCmd).Standalone()
	rootCmd.AddCommand(copyClosureCmd)
}
