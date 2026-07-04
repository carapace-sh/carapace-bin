package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var contentsCmd = &cobra.Command{
	Use:   "contents",
	Short: "List files installed by a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "List dependencies of a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var dependentsCmd = &cobra.Command{
	Use:   "dependents",
	Short: "List installed ports that depend on a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Display the path to the directory containing a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Write arguments after expansion of pseudo-portnames and selectors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var gohomeCmd = &cobra.Command{
	Use:   "gohome",
	Short: "Open the home page for a port in the default browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display meta-information about a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var installedCmd = &cobra.Command{
	Use:   "installed",
	Short: "List installed ports and their activation status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available ports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var livecheckCmd = &cobra.Command{
	Use:   "livecheck",
	Short: "Check if software has been updated since the Portfile was last modified",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "Print the install location of a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Parse and show log files for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var logfileCmd = &cobra.Command{
	Use:   "logfile",
	Short: "Display the path to the log file for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Display notes for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "List installed ports that need an upgrade",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var providesCmd = &cobra.Command{
	Use:   "provides",
	Short: "Determine which port owns a given file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var rdepsCmd = &cobra.Command{
	Use:   "rdeps",
	Short: "Recursively list dependencies of a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var rdependentsCmd = &cobra.Command{
	Use:   "rdependents",
	Short: "Recursively list installed ports that depend on a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for ports by name or description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Display the URL for the path of a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var variantsCmd = &cobra.Command{
	Use:   "variants",
	Short: "List available variants for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the release number of the installed MacPorts infrastructure",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(contentsCmd).Standalone()
	contentsCmd.Flags().Bool("size", false, "Print file sizes")
	contentsCmd.Flags().String("units", "", "Size unit (B, K, Ki, Mi, Gi, kB, MB, GB)")
	rootCmd.AddCommand(contentsCmd)

	carapace.Gen(depsCmd).Standalone()
	depsCmd.Flags().Bool("index", false, "Use cached port index instead of Portfile")
	depsCmd.Flags().Bool("no-build", false, "Exclude build-time dependencies")
	depsCmd.Flags().Bool("no-test", false, "Exclude test dependencies")
	rootCmd.AddCommand(depsCmd)

	carapace.Gen(dependentsCmd).Standalone()
	rootCmd.AddCommand(dependentsCmd)

	carapace.Gen(dirCmd).Standalone()
	rootCmd.AddCommand(dirCmd)

	carapace.Gen(echoCmd).Standalone()
	rootCmd.AddCommand(echoCmd)

	carapace.Gen(gohomeCmd).Standalone()
	rootCmd.AddCommand(gohomeCmd)

	carapace.Gen(infoCmd).Standalone()
	infoCmd.Flags().Bool("category", false, "Display category")
	infoCmd.Flags().Bool("depends", false, "Display dependencies")
	infoCmd.Flags().Bool("fullname", false, "Display full name with version")
	infoCmd.Flags().Bool("homepage", false, "Display homepage")
	infoCmd.Flags().Bool("index", false, "Use cached port index")
	infoCmd.Flags().Bool("line", false, "Output in single-line format")
	infoCmd.Flags().Bool("maintainer", false, "Display maintainer")
	infoCmd.Flags().Bool("pretty", false, "Pretty-print output")
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(installedCmd).Standalone()
	rootCmd.AddCommand(installedCmd)

	carapace.Gen(listCmd).Standalone()
	rootCmd.AddCommand(listCmd)

	carapace.Gen(livecheckCmd).Standalone()
	rootCmd.AddCommand(livecheckCmd)

	carapace.Gen(locationCmd).Standalone()
	rootCmd.AddCommand(locationCmd)

	carapace.Gen(logCmd).Standalone()
	logCmd.Flags().String("phase", "", "Filter log by phase")
	logCmd.Flags().String("verbosity", "", "Filter log by verbosity level")
	rootCmd.AddCommand(logCmd)

	carapace.Gen(logfileCmd).Standalone()
	rootCmd.AddCommand(logfileCmd)

	carapace.Gen(notesCmd).Standalone()
	rootCmd.AddCommand(notesCmd)

	carapace.Gen(outdatedCmd).Standalone()
	rootCmd.AddCommand(outdatedCmd)

	carapace.Gen(providesCmd).Standalone()
	rootCmd.AddCommand(providesCmd)

	carapace.Gen(rdepsCmd).Standalone()
	rdepsCmd.Flags().Bool("full", false, "Show every dependency in full traversal")
	rdepsCmd.Flags().Bool("index", false, "Use cached port index")
	rdepsCmd.Flags().Bool("no-build", false, "Exclude build-time dependencies")
	rdepsCmd.Flags().Bool("no-test", false, "Exclude test dependencies")
	rootCmd.AddCommand(rdepsCmd)

	carapace.Gen(rdependentsCmd).Standalone()
	rdependentsCmd.Flags().Bool("full", false, "Show every dependent in full traversal")
	rootCmd.AddCommand(rdependentsCmd)

	carapace.Gen(searchCmd).Standalone()
	searchCmd.Flags().Bool("case-sensitive", false, "Case-sensitive search")
	searchCmd.Flags().Bool("category", false, "Search in categories")
	searchCmd.Flags().Bool("description", false, "Search in descriptions")
	searchCmd.Flags().Bool("exact", false, "Match literal string exactly")
	searchCmd.Flags().Bool("glob", false, "Treat search string as glob pattern (default)")
	searchCmd.Flags().Bool("homepage", false, "Search in homepages")
	searchCmd.Flags().Bool("line", false, "Output in single-line format")
	searchCmd.Flags().Bool("maintainer", false, "Search by maintainer")
	searchCmd.Flags().Bool("name", false, "Search only port names")
	searchCmd.Flags().Bool("regex", false, "Treat search string as regular expression")
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(urlCmd).Standalone()
	rootCmd.AddCommand(urlCmd)

	carapace.Gen(variantsCmd).Standalone()
	rootCmd.AddCommand(variantsCmd)

	carapace.Gen(versionCmd).Standalone()
	rootCmd.AddCommand(versionCmd)

	carapace.Gen(providesCmd).PositionalCompletion(carapace.ActionFiles())
}
