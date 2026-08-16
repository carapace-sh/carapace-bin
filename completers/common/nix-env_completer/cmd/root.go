package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-env",
	Short: "manage user environments",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-env.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("attr", "A", false, "Interpret args as attribute paths")
	rootCmd.Flags().BoolP("available", "a", false, "Query available packages")
	rootCmd.Flags().Bool("compare-versions", false, "Compare installed vs available versions")
	rootCmd.Flags().Bool("delete-generations", false, "Delete profile generations")
	rootCmd.Flags().Bool("description", false, "Print one-line description")
	rootCmd.Flags().Bool("drv-path", false, "Print store derivation path")
	rootCmd.Flags().Bool("dry-run", false, "Print what would be done without doing it")
	rootCmd.Flags().StringP("file", "f", "", "Path to Nix expression")
	rootCmd.Flags().BoolP("from-expression", "E", false, "Interpret args as Nix expressions")
	rootCmd.Flags().String("from-profile", "", "Copy from another profile")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().BoolP("install", "i", false, "Add packages to user environment")
	rootCmd.Flags().Bool("json", false, "JSON output")
	rootCmd.Flags().Bool("list-generations", false, "List profile generations")
	rootCmd.Flags().Bool("meta", false, "Print all meta-attributes")
	rootCmd.Flags().Bool("no-name", false, "Suppress printing of name attribute")
	rootCmd.Flags().Bool("out-path", false, "Print output path")
	rootCmd.Flags().BoolP("prebuilt-only", "b", false, "Only use derivations with pre-built binaries")
	rootCmd.Flags().BoolP("preserve-installed", "P", false, "Don't remove derivations with matching names")
	rootCmd.Flags().Int("priority", 0, "Set priority of derivations being installed")
	rootCmd.Flags().StringP("profile", "p", "", "Path to profile")
	rootCmd.Flags().BoolP("query", "q", false, "Display information about packages")
	rootCmd.Flags().BoolP("remove-all", "r", false, "Remove all previously installed packages first")
	rootCmd.Flags().Bool("rollback", false, "Switch to previous generation")
	rootCmd.Flags().String("set", "", "Set profile to contain exactly one derivation")
	rootCmd.Flags().Bool("set-flag", false, "Modify meta attributes of installed packages")
	rootCmd.Flags().BoolP("status", "s", false, "Print status")
	rootCmd.Flags().BoolP("switch-generation", "G", false, "Switch to a specific generation")
	rootCmd.Flags().BoolP("switch-profile", "S", false, "Set user environment to given profile")
	rootCmd.Flags().Bool("system", false, "Print system attribute")
	rootCmd.Flags().String("system-filter", "", "Filter derivations by platform")
	rootCmd.Flags().BoolP("uninstall", "e", false, "Remove packages from user environment")
	rootCmd.Flags().BoolP("upgrade", "u", false, "Upgrade packages in user environment")
	rootCmd.Flags().Bool("xml", false, "XML output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(".nix"),
	})
}
