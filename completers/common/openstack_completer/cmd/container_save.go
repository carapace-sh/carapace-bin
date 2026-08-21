package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save container contents locally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_saveCmd).Standalone()

	containerCmd.AddCommand(container_saveCmd)
}
