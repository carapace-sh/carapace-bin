package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mcxquery",
	Short: "Managed Client (MCX) compositor query tool",
	Long:  "https://man.freebsd.org/cgi/man.cgi?mcxquery",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("computer", "", "Computer record")
	rootCmd.Flags().Bool("computerOnly", false, "Ignore values for -user and -group")
	rootCmd.Flags().Bool("forApple", false, "Convenience for bug reports")
	rootCmd.Flags().String("format", "", "Output format")
	rootCmd.Flags().String("group", "", "Group record")
	rootCmd.Flags().String("o", "", "Output file path")
	rootCmd.Flags().Bool("raw", false, "Dump Directory Service data")
	rootCmd.Flags().Bool("useCache", false, "Return cached settings")
	rootCmd.Flags().String("user", "", "User record")
	rootCmd.Flags().Bool("version", false, "Version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("space", "tab", "xml"),
		"o":      carapace.ActionFiles(),
	})
}
