package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getactiveschemeCmd = &cobra.Command{
	Use:   "getactivescheme",
	Short: "retrieve the active power plan",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getactiveschemeCmd).Standalone()
	rootCmd.AddCommand(getactiveschemeCmd)
}
