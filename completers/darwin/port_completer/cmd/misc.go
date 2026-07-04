package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Print the contents of a Portfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Change to the directory containing a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Check for common issues in the user's environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open a Portfile with the default editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Display the path to the Portfile for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display usage information for an action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Verify a Portfile conforms to MacPorts standards",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load a port's daemon via launchctl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Create or update a local mirror of distfiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pkgCmd = &cobra.Command{
	Use:   "pkg",
	Short: "Create a macOS installer package of a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var mpkgCmd = &cobra.Command{
	Use:   "mpkg",
	Short: "Create a macOS installer metapackage of a port and dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var dmgCmd = &cobra.Command{
	Use:   "dmg",
	Short: "Create a disk image containing a macOS package of a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var mdmgCmd = &cobra.Command{
	Use:   "mdmg",
	Short: "Create a disk image containing a metapackage of a port and dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload a port's daemon via launchctl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var unloadCmd = &cobra.Command{
	Use:   "unload",
	Short: "Unload a port's daemon via launchctl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var workCmd = &cobra.Command{
	Use:   "work",
	Short: "Display the path to the work directory for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var distcheckCmd = &cobra.Command{
	Use:   "distcheck",
	Short: "Check if distfiles have not changed and can be fetched",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var distfilesCmd = &cobra.Command{
	Use:   "distfiles",
	Short: "Display each distfile, its checksums, and fetch URLs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "Update checksums and revision for a new version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()
	rootCmd.AddCommand(catCmd)

	carapace.Gen(cdCmd).Standalone()
	rootCmd.AddCommand(cdCmd)

	carapace.Gen(diagnoseCmd).Standalone()
	diagnoseCmd.Flags().Bool("quiet", false, "Only display warnings and errors")
	rootCmd.AddCommand(diagnoseCmd)

	carapace.Gen(editCmd).Standalone()
	editCmd.Flags().String("editor", "", "Specify the editor to use")
	rootCmd.AddCommand(editCmd)

	carapace.Gen(fileCmd).Standalone()
	rootCmd.AddCommand(fileCmd)

	carapace.Gen(helpCmd).Standalone()
	rootCmd.AddCommand(helpCmd)

	carapace.Gen(lintCmd).Standalone()
	lintCmd.Flags().Bool("nitpick", false, "Enable nitpick mode for stricter checks")
	rootCmd.AddCommand(lintCmd)

	carapace.Gen(loadCmd).Standalone()
	rootCmd.AddCommand(loadCmd)

	carapace.Gen(mirrorCmd).Standalone()
	mirrorCmd.Flags().Bool("new", false, "Only mirror distfiles not already present")
	rootCmd.AddCommand(mirrorCmd)

	carapace.Gen(pkgCmd).Standalone()
	rootCmd.AddCommand(pkgCmd)

	carapace.Gen(mpkgCmd).Standalone()
	rootCmd.AddCommand(mpkgCmd)

	carapace.Gen(dmgCmd).Standalone()
	rootCmd.AddCommand(dmgCmd)

	carapace.Gen(mdmgCmd).Standalone()
	rootCmd.AddCommand(mdmgCmd)

	carapace.Gen(reloadCmd).Standalone()
	rootCmd.AddCommand(reloadCmd)

	carapace.Gen(unloadCmd).Standalone()
	rootCmd.AddCommand(unloadCmd)

	carapace.Gen(workCmd).Standalone()
	rootCmd.AddCommand(workCmd)

	carapace.Gen(distcheckCmd).Standalone()
	rootCmd.AddCommand(distcheckCmd)

	carapace.Gen(distfilesCmd).Standalone()
	rootCmd.AddCommand(distfilesCmd)

	carapace.Gen(bumpCmd).Standalone()
	rootCmd.AddCommand(bumpCmd)
}
