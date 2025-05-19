package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCompletionsCmd = &cobra.Command{
	Use:   "install-completions",
	Short: "Install shell completions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCompletionsCmd).Standalone()

	installCompletionsCmd.Flags().Bool("uninstall", false, "")
	rootCmd.AddCommand(installCompletionsCmd)
}
