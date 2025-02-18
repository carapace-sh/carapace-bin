package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var prependCmd = &cobra.Command{
	Use:   "prepend",
	Short: "prepend RULE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prependCmd)

	rootCmd.AddCommand(prependCmd)
}
