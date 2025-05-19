package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [flags] COMMAND",
	Short: "Use the provisioner to cleanup any changes made to external systems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd)

	initCmd.Flags().Bool("help", false, "Show help message and exit")

	rootCmd.AddCommand(initCmd)
}
