package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "print OpenSSL version information",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd)

	versionCmd.Flags().BoolS("a", "a", false, "Show all data")
	versionCmd.Flags().BoolS("b", "b", false, "Show build date")
	versionCmd.Flags().BoolS("c", "c", false, "Show CPU settings info")
	versionCmd.Flags().BoolS("d", "d", false, "Show configuration directory")
	versionCmd.Flags().BoolS("e", "e", false, "Show engines directory")
	versionCmd.Flags().BoolS("f", "f", false, "Show compiler flags used")
	versionCmd.Flags().BoolS("m", "m", false, "Show modules directory")
	versionCmd.Flags().BoolS("o", "o", false, "Show some internal datatype options")
	versionCmd.Flags().BoolS("p", "p", false, "Show target build platform")
	versionCmd.Flags().BoolS("r", "r", false, "Show random seeding options")
	versionCmd.Flags().BoolS("v", "v", false, "Show library version")
	versionCmd.Flags().BoolS("w", "w", false, "Show Windows install context")
	rootCmd.AddCommand(versionCmd)
}
