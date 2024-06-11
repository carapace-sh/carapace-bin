package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag [OPTIONS] <COMMAND>",
	Short: "Manage Tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tagCmd).Standalone()

	rootCmd.AddCommand(tagCmd)
}
