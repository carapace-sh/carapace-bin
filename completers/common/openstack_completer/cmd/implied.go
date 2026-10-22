package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var impliedCmd = &cobra.Command{
	Use:   "implied",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(impliedCmd).Standalone()

	rootCmd.AddCommand(impliedCmd)
}
