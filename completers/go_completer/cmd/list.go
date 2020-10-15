package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list packages or modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().StringS("f", "f", "", "specify an alternate format for the list")
	listCmd.Flags().BoolS("m", "m", false, "list modules instead of packages")
	listCmd.Flags().BoolS("u", "u", false, "add information about available upgrades")
	listCmd.Flags().BoolS("e", "e", false, "change the handling of erroneous packages")
	listCmd.Flags().Bool("compiled", false, "set CompiledGoFiles to the Go source files presented to the compiler")
	listCmd.Flags().Bool("deps", false, "iterate over not just the named packages but also all their dependencies")
	listCmd.Flags().Bool("export", false, "set the Export field to the name of a file containing up-to-date export information")
	listCmd.Flags().Bool("find", false, "identify the named packages but not resolve their dependencies")
	listCmd.Flags().Bool("json", false, "print package data in JSON format")
	listCmd.Flags().Bool("test", false, "report not only the named packages but also their test binaries")
	listCmd.Flags().Bool("versions", false, "set the Module's Versions field to list of all known versions")
	rootCmd.AddCommand(listCmd)
}
