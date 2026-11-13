package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var introspectCmd = &cobra.Command{
	Use:   "introspect",
	Short: "Introspect project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(introspectCmd).Standalone()

	introspectCmd.Flags().BoolP("all", "a", false, "Print all available information.")
	introspectCmd.Flags().Bool("ast", false, "Dump the AST of the meson file.")
	introspectCmd.Flags().String("backend", "", "The backend to use for the --buildoptions introspection.")
	introspectCmd.Flags().Bool("benchmarks", false, "List all benchmarks.")
	introspectCmd.Flags().Bool("buildoptions", false, "List all build options.")
	introspectCmd.Flags().Bool("buildsystem-files", false, "List files that make up the build system.")
	introspectCmd.Flags().Bool("compilers", false, "List used compilers.")
	introspectCmd.Flags().Bool("dependencies", false, "List external dependencies.")
	introspectCmd.Flags().BoolP("force-object-output", "f", false, "Always use the new JSON format for multiple entries (even for 0 and 1 introspection commands)")
	introspectCmd.Flags().BoolP("indent", "i", false, "Enable pretty printed JSON.")
	introspectCmd.Flags().Bool("install-plan", false, "List all installed files and directories with their details.")
	introspectCmd.Flags().Bool("installed", false, "List all installed files and directories.")
	introspectCmd.Flags().Bool("machines", false, "Information about host, build, and target machines.")
	introspectCmd.Flags().Bool("projectinfo", false, "Information about projects.")
	introspectCmd.Flags().Bool("scan-dependencies", false, "Scan for dependencies used in the meson.build file.")
	introspectCmd.Flags().Bool("targets", false, "List top level targets.")
	introspectCmd.Flags().Bool("tests", false, "List all unit tests.")
	applyCommonFlagCompletion(introspectCmd)
	rootCmd.AddCommand(introspectCmd)
}
