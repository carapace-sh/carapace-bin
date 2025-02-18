package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var denyCmd = &cobra.Command{
	Use:   "deny",
	Short: "add deny rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(denyCmd)

	rootCmd.AddCommand(denyCmd)
}
