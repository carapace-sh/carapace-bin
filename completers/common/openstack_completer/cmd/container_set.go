package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set container properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_setCmd).Standalone()

	container_setCmd.Flags().String("property", "", "Set a property on this container (repeat option to set multiple properties)")
	container_setCmd.MarkFlagRequired("property")
	containerCmd.AddCommand(container_setCmd)
}
