package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var object_storeCmd = &cobra.Command{
	Use:   "store",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_storeCmd).Standalone()

	objectCmd.AddCommand(object_storeCmd)
}
