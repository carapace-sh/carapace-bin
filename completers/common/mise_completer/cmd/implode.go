package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var implodeCmd = &cobra.Command{
	Use:   "implode",
	Short: "Removes mise CLI and all related data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(implodeCmd).Standalone()

	implodeCmd.Flags().Bool("config", false, "Also remove config files")
	implodeCmd.Flags().BoolP("dry-run", "n", false, "Show what would be removed without actually removing")
	rootCmd.AddCommand(implodeCmd)
}
