package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "newfs_apfs",
	Short: "Create APFS filesystem",
	Long:  "https://keith.github.io/xcode-manpages/newfs_apfs.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Add a volume to an existing container")
	rootCmd.Flags().BoolS("C", "C", false, "Create a container only, with no volume")
	rootCmd.Flags().BoolS("D", "D", false, "Opt in of UUID from role")
	rootCmd.Flags().BoolS("E", "E", false, "Enable volume encryption")
	rootCmd.Flags().BoolS("T", "T", false, "Resulting file system is case-sensitive")
	rootCmd.Flags().BoolS("a", "a", false, "Add all existing volumes")
	rootCmd.Flags().BoolS("c", "c", false, "Create a case-sensitive volume")
	rootCmd.Flags().BoolS("e", "e", false, "Create a case-sensitive volume")
	rootCmd.Flags().BoolS("i", "i", false, "Create a case-insensitive volume")
	rootCmd.Flags().BoolS("w", "w", false, "Wait for volume to be fully published")
	rootCmd.Flags().BoolS("x", "x", false, "Do not warn about unusual volume names")

	rootCmd.Flags().StringS("G", "G", "", "GID of the root volume")
	rootCmd.Flags().StringS("R", "R", "", "Volume role")
	rootCmd.Flags().StringS("S", "S", "", "Password key for volume encryption")
	rootCmd.Flags().StringS("U", "U", "", "UID of the root volume")
	rootCmd.Flags().StringS("b", "b", "", "Block size of the container")
	rootCmd.Flags().StringS("o", "o", "", "Additional volume formatting options")
	rootCmd.Flags().StringS("q", "q", "", "Volume quota (upper limit)")
	rootCmd.Flags().StringS("r", "r", "", "Volume reserve size")
	rootCmd.Flags().StringS("s", "s", "", "Volume size")
	rootCmd.Flags().StringS("v", "v", "", "Volume name")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"G": carapace.ActionValues(),
		"R": carapace.ActionValues("none", "system", "data", "preboot", "recovery", "vm"),
		"S": carapace.ActionValues(),
		"U": carapace.ActionValues(),
		"b": carapace.ActionValues("4096", "8192", "16384", "32768", "65536", "131072", "262144", "524288", "1048576"),
		"o": carapace.ActionValues(),
		"v": carapace.ActionValues(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
