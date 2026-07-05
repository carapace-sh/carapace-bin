package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setactiveCmd = &cobra.Command{
	Use:   "setactive",
	Short: "make a power plan active",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setactiveCmd).Standalone()
	rootCmd.AddCommand(setactiveCmd)
}
