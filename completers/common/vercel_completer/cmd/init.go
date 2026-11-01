package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize example Vercel Projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("force", false, "Force overwrite existing files")

	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionDirectories(),
	)
}
