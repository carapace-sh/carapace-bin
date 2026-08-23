package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "otool",
	Short: "Object file display tool",
	Long:  "https://keith.github.io/xcode-manpages/otool.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("D", "D", false, "Display shared library id name")
	rootCmd.Flags().BoolS("L", "L", false, "Display shared libraries used")
	rootCmd.Flags().BoolS("a", "a", false, "Display archive header")
	rootCmd.Flags().BoolS("d", "d", false, "Display contents of __DATA,__data section")
	rootCmd.Flags().BoolS("f", "f", false, "Display universal headers")
	rootCmd.Flags().BoolS("h", "h", false, "Display Mach-O header")
	rootCmd.Flags().BoolS("l", "l", false, "Display load commands")
	rootCmd.Flags().BoolS("o", "o", false, "Display contents of __OBJC segment")
	rootCmd.Flags().BoolS("r", "r", false, "Display relocation entries")
	rootCmd.Flags().BoolS("t", "t", false, "Display contents of (__TEXT,__text) section")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode (disassemble with -t)")
	rootCmd.Flags().Bool("version", false, "Display the version")

	rootCmd.Flags().String("arch", "", "Specify architecture type")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"arch": carapace.ActionValues("arm64", "arm64e", "x86_64", "x86_64h"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
