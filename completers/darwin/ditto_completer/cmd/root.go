package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ditto",
	Short: "copy directory hierarchies, create and extract archives",
	Long:  "https://keith.github.io/xcode-manpages/ditto.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("acl", false, "Preserve Access Control Lists (ACLs)")
	rootCmd.Flags().String("arch", "", "Thin Universal binaries to the specified architecture")
	rootCmd.Flags().String("bom", "", "Copy only files present in the specified BOM")
	rootCmd.Flags().BoolP("bzip2", "j", false, "Create compressed CPIO archives using bzip2")
	rootCmd.Flags().Bool("clone", false, "Attempt to clone regular files when copying")
	rootCmd.Flags().BoolP("create", "c", false, "Create an archive at the destination path")
	rootCmd.Flags().Bool("extattr", false, "Preserve extended attributes")
	rootCmd.Flags().BoolP("extract", "x", false, "Extract archives")
	rootCmd.Flags().BoolP("gzip", "z", false, "Create compressed CPIO archives using gzip")
	rootCmd.Flags().Bool("help", false, "Print full usage")
	rootCmd.Flags().Bool("hfsCompression", false, "Compress files with filesystem compression")
	rootCmd.Flags().Bool("keepBinaries", false, "Set aside original Mach-O binaries when replacing")
	rootCmd.Flags().String("keepBinariesList", "", "Record kept binary locations in the file at path")
	rootCmd.Flags().String("keepBinariesPattern", "", "Keep files matching the specified regex")
	rootCmd.Flags().Bool("keepParent", false, "Embed the parent directory name src in dst_archive")
	rootCmd.Flags().String("lang", "", "Specify language variants to filter from index bom")
	rootCmd.Flags().BoolP("noCrossDevice", "X", false, "Do not descend into directories on a different device")
	rootCmd.Flags().Bool("noacl", false, "Do not preserve ACLs")
	rootCmd.Flags().Bool("nocache", false, "Do not perform copies using the Unified Buffer Cache")
	rootCmd.Flags().Bool("noclone", false, "Do not attempt to clone files")
	rootCmd.Flags().Bool("noextattr", false, "Do not preserve extended attributes")
	rootCmd.Flags().Bool("nohfsCompression", false, "Do not compress files with filesystem compression")
	rootCmd.Flags().Bool("nonAtomicCopies", false, "Do not perform atomic copies")
	rootCmd.Flags().Bool("nopersistRootless", false, "Do not persist rootless flag/attribute")
	rootCmd.Flags().Bool("nopreserveHFSCompression", false, "Do not preserve filesystem compression")
	rootCmd.Flags().Bool("noqtn", false, "Do not preserve quarantine information")
	rootCmd.Flags().Bool("norsrc", false, "Do not preserve resource forks and HFS meta-data")
	rootCmd.Flags().String("option", "", "Specify an arbitrary key value pair for the copier")
	rootCmd.Flags().String("outBom", "", "Specify an explicit path for the output bom")
	rootCmd.Flags().Bool("password", false, "Prompt for a password to extract encrypted ZIP archive")
	rootCmd.Flags().Bool("persistRootless", false, "Retain SF_RESTRICTED flag or com.apple.rootless")
	rootCmd.Flags().Bool("preserveHFSCompression", false, "Preserve filesystem compression of source files")
	rootCmd.Flags().Bool("qtn", false, "Preserve quarantine information")
	rootCmd.Flags().Bool("rsrc", false, "Preserve resource forks and HFS meta-data")
	rootCmd.Flags().Bool("segmentLargeFiles", false, "Segment files larger than 8GB into multiple entries")
	rootCmd.Flags().Bool("sequesterRsrc", false, "Preserve resource forks in __MACOSX subdirectory")
	rootCmd.Flags().BoolP("verbose", "v", false, "Print a line per source directory copied")
	rootCmd.Flags().BoolP("veryverbose", "V", false, "Print a line per file copied")
	rootCmd.Flags().Bool("zip", false, "Create or extract from a PKZip archive")
	rootCmd.Flags().String("zlibCompressionLevel", "", "Set the compression level (0-9) for PKZip")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"arch":                carapace.ActionValues("arm64", "x86_64"),
		"bom":                 carapace.ActionFiles(),
		"keepBinariesList":    carapace.ActionFiles(),
		"keepBinariesPattern": carapace.ActionFiles(),
		"outBom":              carapace.ActionFiles(),
		"zlibCompressionLevel": carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8", "9"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}