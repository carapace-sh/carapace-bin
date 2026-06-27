package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Perform swiftly initialization into your user account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().BoolP("help", "h", false, "Show help information")
	initCmd.Flags().BoolP("no-modify-profile", "n", false, "Do not modify the profile file to set environment variables")
	initCmd.Flags().BoolP("overwrite", "o", false, "Overwrite the existing swiftly installation")
	initCmd.Flags().String("platform", "", "Specify the current Linux platform for swiftly")
	initCmd.Flags().Bool("quiet-shell-followup", false, "Quiet shell follow up commands")
	initCmd.Flags().Bool("skip-install", false, "Skip installing the latest toolchain")

	rootCmd.AddCommand(initCmd)
}
