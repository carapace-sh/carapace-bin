package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var npm_tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "manage tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(npm_tagCmd).Standalone()

	npmCmd.AddCommand(npm_tagCmd)
}
