package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keys_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Output the path to the keys.json file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keys_pathCmd).Standalone()

	keysCmd.AddCommand(keys_pathCmd)
}
