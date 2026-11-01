package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configurePendingCmd = &cobra.Command{
	Use:     "configure-pending",
	Aliases: []string{"cp"},
	Short:   "perform any system configuration if any packages are in a pending state",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configurePendingCmd).Standalone()

	configurePendingCmd.Flags().BoolP("dry-run", "n", false, "only show what would happen, do not actually perform changes")
	configurePendingCmd.Flags().Bool("ignore-comar", false, "bypass system configuration")
	configurePendingCmd.Flags().Bool("ignore-dependency", false, "do not attempt to validate dependencies")
	configurePendingCmd.Flags().Bool("ignore-safety", false, "ignore safety switch on system.base component")

	rootCmd.AddCommand(configurePendingCmd)
}
