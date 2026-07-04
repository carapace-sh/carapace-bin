package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "repair_packages",
	Short: "Verify or repair filesystem permissions and flags for installed packages",
	Long:  "https://www.manpagez.com/man/8/repair_packages/",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("list-standard-pkgs", false, "output the list of standard package-id strings")
	rootCmd.Flags().Bool("no-fscompression", false, "don't use filesystem compression (undocumented)")
	rootCmd.Flags().String("output-format", "", "use a special output format (selected by a number)")
	rootCmd.Flags().String("pkg", "", "add a specific package by its package identifier")
	rootCmd.Flags().Bool("repair", false, "repair the permissions/flags of the specified package(s)")
	rootCmd.Flags().Bool("standard-pkgs", false, "add the standard built-in Apple packages to the list")
	rootCmd.Flags().Bool("verify", false, "verify the permissions/flags of the specified package(s)")
	rootCmd.Flags().String("volume", "", "perform all operations on the specified volume")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValues("0", "1", "2"),
		"volume":        carapace.ActionDirectories(),
	})
}
