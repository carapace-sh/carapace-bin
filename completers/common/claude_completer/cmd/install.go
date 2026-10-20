package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Claude Code native build",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("force", false, "Force installation even if already installed")
	installCmd.Flags().BoolP("help", "h", false, "Display help for command")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalCompletion(
		carapace.ActionValues("stable", "latest"),
	)
}
