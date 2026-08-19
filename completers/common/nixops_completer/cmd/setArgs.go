package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setArgsCmd = &cobra.Command{
	Use:   "set-args",
	Short: "persistently set deployment arguments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setArgsCmd).Standalone()
	rootCmd.AddCommand(setArgsCmd)
}
