package cmd

import (
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chmod",
	Short: "change file modes",
	Long:  "https://keith.github.io/xcode-manpages/chmod.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	for i, arg := range os.Args {
		if !strings.HasPrefix(arg, "-") || len(arg) < 2 {
			continue
		}
		if arg == "--" {
			break
		}
		switch arg[1] {
		case 'X', 'r', 's', 't', 'w', 'x':
			switch i {
			case len(os.Args) - 1:
				continue
			default:
				os.Args[i] = "-=" + arg[1:]
			}
		}
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "Compute and display canonical ACL for file")
	rootCmd.Flags().BoolS("E", "E", false, "Read ACL information from stdin")
	rootCmd.Flags().BoolS("H", "H", false, "If the file is a symbolic link, change the mode of the link itself (with -R)")
	rootCmd.Flags().BoolS("L", "L", false, "Traverse all symbolic links to directories (with -R)")
	rootCmd.Flags().BoolS("N", "N", false, "Clear all ACL entries for the file")
	rootCmd.Flags().BoolS("P", "P", false, "Do not traverse any symbolic links to directories (with -R)")
	rootCmd.Flags().BoolS("R", "R", false, "Recursively change file mode bits")
	rootCmd.Flags().BoolS("f", "f", false, "Do not display a diagnostic message if chmod could not modify the mode")
	rootCmd.Flags().BoolS("h", "h", false, "If the file is a symbolic link, change the mode of the link itself")
	rootCmd.Flags().BoolS("v", "v", false, "Cause chmod to be verbose, showing files as the mode is modified")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
