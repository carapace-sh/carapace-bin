package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pass"
	"github.com/spf13/cobra"
)

var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "rename or move old-path to new-path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mvCmd).Standalone()
	mvCmd.Flags().BoolP("force", "f", false, "mv forcefully")

	rootCmd.AddCommand(mvCmd)

	carapace.Gen(mvCmd).PositionalCompletion(
		pass.ActionPasswords(),
		pass.ActionPasswords(),
	)
}
