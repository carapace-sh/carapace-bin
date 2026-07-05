package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorychipCmd = &cobra.Command{
	Use:   "memorychip",
	Short: "memory chip management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorychipCmd).Standalone()
	rootCmd.AddCommand(memorychipCmd)
}
