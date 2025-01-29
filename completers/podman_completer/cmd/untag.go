package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var untagCmd = &cobra.Command{
	Use:   "untag IMAGE [IMAGE...]",
	Short: "Remove a name from a local image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(untagCmd).Standalone()

	rootCmd.AddCommand(untagCmd)
}
