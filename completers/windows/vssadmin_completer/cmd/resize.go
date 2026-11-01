package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "resize shadow copy storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resizeCmd).Standalone()
	rootCmd.AddCommand(resizeCmd)
}
