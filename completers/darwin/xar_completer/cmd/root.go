package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xar",
	Short: "eXtensible ARchiver",
	Long:  "https://keith.github.io/xcode-manpages/xar.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("C", "", "Extract to the specified path")
	rootCmd.Flags().Bool("P", false, "Set ownership based on uid/gid")
	rootCmd.Flags().Bool("a", false, "Synonym for --compression=lzma")
	rootCmd.Flags().String("compression", "", "Specify compression type")
	rootCmd.Flags().String("compression-args", "", "Specify arguments to the compression engine")
	rootCmd.Flags().BoolP("create", "c", false, "Create an archive")
	rootCmd.Flags().Bool("dump-header", false, "Print the xar binary header information")
	rootCmd.Flags().String("dump-toc", "", "Dump the XML header into the specified file")
	rootCmd.Flags().Bool("dump-toc-cksum", false, "Dump the ToC checksum to stdout")
	rootCmd.Flags().BoolP("extract", "x", false, "Extract an archive")
	rootCmd.Flags().String("extract-subdoc", "", "Extract the specified subdocument")
	rootCmd.Flags().StringP("file", "f", "", "Filename to use for creation, listing or extraction")
	rootCmd.Flags().String("file-cksum", "", "Specify the hashing algorithm for file content verification")
	rootCmd.Flags().Bool("j", false, "Synonym for --compression=bzip2")
	rootCmd.Flags().Bool("l", false, "Stay on the local device")
	rootCmd.Flags().BoolP("list", "t", false, "List the contents of an archive")
	rootCmd.Flags().Bool("list-subdocs", false, "List the subdocuments in the XML header")
	rootCmd.Flags().Bool("p", false, "Set ownership based on symbolic names")
	rootCmd.Flags().String("s", "", "Specify the file to extract subdocuments to")
	rootCmd.Flags().String("toc-cksum", "", "Specify the hashing algorithm for ToC verification")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.Flags().Bool("z", false, "Synonym for --compression=gzip")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"compression":      carapace.ActionValues("none", "gzip", "bzip2", "lzma"),
		"compression-args": carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8", "9"),
		"dump-toc":         carapace.ActionFiles(),
		"extract-subdoc":   carapace.ActionFiles(),
		"file":             carapace.ActionFiles(),
		"file-cksum":       carapace.ActionValues("sha1", "sha256", "sha512", "md5"),
		"s":                carapace.ActionFiles(),
		"toc-cksum":        carapace.ActionValues("sha1", "sha256", "sha512", "md5"),
	})
}
