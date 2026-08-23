package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pkgutil",
	Short: "query and manipulate macOS Installer packages and receipts",
	Long:  "https://keith.github.io/xcode-manpages/pkgutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "A brief summary of commands and usage")
	rootCmd.Flags().BoolP("force", "f", false, "Don't ask for confirmation before performing destructive operation")
	rootCmd.Flags().BoolP("verbose", "v", false, "Output in human-readable format")
	rootCmd.Flags().String("volume", "", "Perform all operations on the specified volume")
	rootCmd.Flags().String("edit-pkg", "", "Specifies an existing receipt to be modified in-place by --learn")
	rootCmd.Flags().Bool("only-files", false, "List only files in --files listing")
	rootCmd.Flags().Bool("only-dirs", false, "List only directories in --files listing")
	rootCmd.Flags().Bool("regexp", false, "Try to match package-id arguments as a regular expression")
	rootCmd.Flags().String("pkgs", "", "List all installed package IDs matching REGEXP")
	rootCmd.Flags().Bool("pkgs-plist", false, "List all installed package IDs in plist format")
	rootCmd.Flags().Bool("packages", false, "List all installed package IDs")
	rootCmd.Flags().Bool("groups", false, "List all of the package groups")
	rootCmd.Flags().Bool("groups-plist", false, "List all package groups in plist format")
	rootCmd.Flags().String("files", "", "List all files installed under the package-id")
	rootCmd.Flags().String("export-plist", "", "Print receipt information in plist format")
	rootCmd.Flags().String("pkg-info", "", "Print extended information about the specified package-id")
	rootCmd.Flags().String("pkg-info-plist", "", "Print extended info in plist format")
	rootCmd.Flags().String("forget", "", "Discard all receipt data about package-id")
	rootCmd.Flags().String("learn", "", "Update the ACLs of the given path in the receipt")
	rootCmd.Flags().String("pkg-groups", "", "List package groups this package-id is a member of")
	rootCmd.Flags().String("group-pkgs", "", "List packages that are members of this group-id")
	rootCmd.Flags().String("file-info", "", "Show the metadata known about path")
	rootCmd.Flags().String("file-info-plist", "", "Show metadata in plist format")
	rootCmd.Flags().String("expand", "", "Expand the flat package at pkg-path into a directory")
	rootCmd.Flags().String("flatten", "", "Flatten the directory into a flat package")
	rootCmd.Flags().String("bom", "", "Extract BOM files from the flat pkg at path")
	rootCmd.Flags().String("payload-files", "", "List files archived within the payload")
	rootCmd.Flags().String("check-signature", "", "Check the signature on the package at path")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bom":              carapace.ActionFiles(),
		"expand":           carapace.ActionFiles(),
		"file-info":        carapace.ActionFiles(),
		"file-info-plist":  carapace.ActionFiles(),
		"flatten":          carapace.ActionFiles(),
		"learn":            carapace.ActionFiles(),
		"payload-files":    carapace.ActionFiles(),
		"check-signature":  carapace.ActionFiles(),
		"volume":           carapace.ActionFiles(),
	})
}