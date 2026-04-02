package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configurePendingCmd = &cobra.Command{
	Use:     "configure-pending",
	Short:   "Configure pending packages",
	Aliases: []string{"cp"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configurePendingCmd).Standalone()

	configurePendingCmd.Flags().Bool("ignore-dependency", false, "Do not take dependency info into account")
	configurePendingCmd.Flags().Bool("ignore-comar", false, "Bypass COMAR configuration agent")
	configurePendingCmd.Flags().Bool("ignore-safety", false, "Bypass safety switch")
	configurePendingCmd.Flags().BoolP("dry-run", "n", false, "Do not perform any action, just show what would be done")

	rootCmd.AddCommand(configurePendingCmd)
}
