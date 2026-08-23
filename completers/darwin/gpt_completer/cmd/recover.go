package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "recover the GPT",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recoverCmd).Standalone()
	rootCmd.AddCommand(recoverCmd)
}