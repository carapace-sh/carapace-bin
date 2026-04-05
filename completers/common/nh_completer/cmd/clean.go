package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Enhanced nix cleanup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	rootCmd.AddCommand(cleanCmd)
}

func addCleanFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("keep", "k", "1", "At least keep this number of generations")
	cmd.Flags().StringP("keep-since", "K", "0h", "At least keep gcroots and generations in this time range since now")
	cmd.Flags().BoolP("dry", "n", false, "Only print actions, without performing them")
	cmd.Flags().BoolP("ask", "a", false, "Ask for confirmation")
	cmd.Flags().Bool("no-gc", false, "Don't run nix store --gc")
	cmd.Flags().Bool("no-gcroots", false, "Don't clean gcroots")
	cmd.Flags().Bool("optimise", false, "Run nix-store --optimise after gc")
	cmd.Flags().String("max", "", "Pass --max to nix store gc")
}
