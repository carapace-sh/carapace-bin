package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pass"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove existing password or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()
	rmCmd.Flags().BoolP("force", "f", false, "remove forcefully")
	rmCmd.Flags().BoolP("recursive", "r", false, "remove recursively")

	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalCompletion(
		pass.ActionPasswords(),
	)
}
