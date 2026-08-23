package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ndp",
	Short: "control/diagnose IPv6 neighbor discovery protocol",
	Long:  "https://keith.github.io/xcode-manpages/ndp.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("A", "A", "", "Wait for resolution")
	rootCmd.Flags().BoolS("H", "H", false, "Print ND protocol header")
	rootCmd.Flags().StringS("I", "I", "", "Default interface")
	rootCmd.Flags().BoolS("P", "P", false, "Print prefix list")
	rootCmd.Flags().BoolS("R", "R", false, "Print router list")
	rootCmd.Flags().StringS("W", "W", "", "Config file")
	rootCmd.Flags().BoolS("a", "a", false, "Dump NDP entries")
	rootCmd.Flags().BoolS("c", "c", false, "Clear NDP statistics")
	rootCmd.Flags().StringS("d", "d", "", "Delete NDP entry")
	rootCmd.Flags().StringS("f", "f", "", "Flush NDP entries")
	rootCmd.Flags().StringS("i", "i", "", "Interface")
	rootCmd.Flags().BoolS("l", "l", false, "Long output")
	rootCmd.Flags().BoolS("n", "n", false, "Numeric output")
	rootCmd.Flags().BoolS("p", "p", false, "Print default router list")
	rootCmd.Flags().BoolS("r", "r", false, "Print routing table")
	rootCmd.Flags().StringS("s", "s", "", "Set NDP entry")
	rootCmd.Flags().BoolS("t", "t", false, "Timeout")
	rootCmd.Flags().BoolS("w", "w", false, "Wait")
	rootCmd.Flags().BoolS("x", "x", false, "Extended output")
	rootCmd.Flags().BoolS("z", "z", false, "Zero NDP statistics")
}
