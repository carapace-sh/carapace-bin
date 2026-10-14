package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Runs util (default: env) inside the environment used for command shim execution",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()

	rootCmd.AddCommand(envCmd)
}
