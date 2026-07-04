package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ioreg",
	Short: "display IOKit registry",
	Long:  "https://keith.github.io/xcode-manpages/ioreg.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("archive", "a", false, "Archive the output in XML")
	rootCmd.Flags().BoolP("bold", "b", false, "Show the object name in bold")
	rootCmd.Flags().StringP("class", "c", "", "Show properties only for objects of the specified class")
	rootCmd.Flags().StringP("depth", "d", "", "Limit tree traversal to the specified depth")
	rootCmd.Flags().BoolP("format", "f", false, "Enable smart formatting")
	rootCmd.Flags().BoolP("hex", "x", false, "Show data and numbers as hexadecimal")
	rootCmd.Flags().BoolP("inheritance", "i", false, "Show the object inheritance")
	rootCmd.Flags().StringP("key", "k", "", "Show properties only if the object has the specified key")
	rootCmd.Flags().BoolP("location", "t", false, "Show tree location of each subtree")
	rootCmd.Flags().StringP("name", "n", "", "Show properties only if the object has the specified name")
	rootCmd.Flags().StringP("plane", "p", "", "Traverse the registry over the specified plane")
	rootCmd.Flags().BoolP("properties", "l", false, "Show properties for all displayed objects")
	rootCmd.Flags().BoolP("subtrees", "r", false, "Show subtrees rooted by objects matching criteria")
	rootCmd.Flags().StringP("width", "w", "", "Clip the output to the specified line width")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"plane": carapace.ActionValues("IODeviceTree", "IOService", "IOAudio", "IOPlatformExpertDevice"),
	})
}
