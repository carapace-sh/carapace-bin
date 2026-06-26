package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kubercCmd = &cobra.Command{
	Use:   "kuberc SUBCOMMAND",
	Short: "Manage kuberc configuration files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubercCmd).Standalone()

	rootCmd.AddCommand(kubercCmd)
}
