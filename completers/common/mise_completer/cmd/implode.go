package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var implodeCmd = &cobra.Command{
	Use:   "implode",
	Short: "Remove the mise CLI and all related data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(implodeCmd).Standalone()

	implodeCmd.Flags().Bool("config", false, "Also remove the config directory")
	implodeCmd.Flags().BoolP("dry-run", "n", false, "Print what would be removed, without removing anything")
	implodeCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(implodeCmd)
}
