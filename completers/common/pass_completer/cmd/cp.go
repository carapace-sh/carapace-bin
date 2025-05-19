package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pass"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "copies old-path to new-path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cpCmd).Standalone()
	cpCmd.Flags().BoolP("force", "f", false, "cp forcefully")

	rootCmd.AddCommand(cpCmd)

	carapace.Gen(cpCmd).PositionalCompletion(
		pass.ActionPasswords(),
		pass.ActionPasswords(),
	)
}
