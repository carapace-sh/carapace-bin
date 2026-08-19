package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-collect-garbage",
	Short: "delete unreachable store objects",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-collect-garbage.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("delete-old", "d", false, "Delete all old generations of profiles")
	rootCmd.Flags().String("delete-older-than", "", "Delete all generations of profiles older than the specified amount")
	rootCmd.Flags().Bool("dry-run", false, "Print what would be deleted without actually deleting")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().String("max-freed", "", "Keep deleting paths until at least bytes have been deleted")
}
