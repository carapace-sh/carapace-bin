package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cached_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cached_imageCmd).Standalone()

	cachedCmd.AddCommand(cached_imageCmd)
}
