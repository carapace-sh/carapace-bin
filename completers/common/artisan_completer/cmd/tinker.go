package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tinkerCmd = &cobra.Command{
	Use:   "tinker [--execute [EXECUTE]] [--] [<include>...]",
	Short: "Interact with your application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tinkerCmd).Standalone()

	tinkerCmd.Flags().String("execute", "", "Execute the given code using Tinker")
	rootCmd.AddCommand(tinkerCmd)
}
