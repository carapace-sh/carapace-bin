package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debuguigetpassCmd = &cobra.Command{
	Use:    "debuguigetpass",
	Short:  "show prompt to type password",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuguigetpassCmd).Standalone()

	debuguigetpassCmd.Flags().StringP("prompt", "p", "", "prompt text")
	rootCmd.AddCommand(debuguigetpassCmd)
}
