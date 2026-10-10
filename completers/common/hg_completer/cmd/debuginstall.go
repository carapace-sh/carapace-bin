package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debuginstallCmd = &cobra.Command{
	Use:    "debuginstall",
	Short:  "test Mercurial installation",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuginstallCmd).Standalone()

	debuginstallCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debuginstallCmd)
}
