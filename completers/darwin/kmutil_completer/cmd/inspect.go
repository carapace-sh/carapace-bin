package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect and display a kext collection's contents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()

	inspectCmd.Flags().StringP("arch", "a", "", "Architecture to use")
	inspectCmd.Flags().StringP("auxiliary-path", "A", "", "Path to the auxiliary kernel collection")
	inspectCmd.Flags().StringP("boot-path", "B", "", "Path to the boot kernel collection")
	inspectCmd.Flags().StringP("bundle-identifier", "b", "", "Search for and include this identifier")
	inspectCmd.Flags().StringP("bundle-path", "p", "", "Include the bundle specified at this path")
	inspectCmd.Flags().String("elide-identifier", "", "Do not include this identifier in the results")
	inspectCmd.Flags().StringP("filter", "f", "", "Filter results by predicate (overridden)")
	inspectCmd.Flags().StringP("filter-all", "F", "", "Filter results by predicate (not overridden)")
	inspectCmd.Flags().BoolP("help", "h", false, "Show help information")
	inspectCmd.Flags().Bool("json", false, "Print the section layout as JSON")
	inspectCmd.Flags().BoolP("no-authorization", "z", false, "Skip approval checks")
	inspectCmd.Flags().Bool("no-header", false, "Don't show the header describing the collection")
	inspectCmd.Flags().Bool("no-variant-extension", false, "Don't append build variant as extension to path")
	inspectCmd.Flags().String("optional-identifier", "", "Search for and include this identifier if possible")
	inspectCmd.Flags().Bool("show-collection-metadata", false, "Print the metadata of the collection(s)")
	inspectCmd.Flags().Bool("show-extension-info", false, "Show extension information (default)")
	inspectCmd.Flags().Bool("show-fileset-entries", false, "Print fileset entries of the mach header(s)")
	inspectCmd.Flags().Bool("show-kernel-uuid", false, "Print the UUID of the kernel in the collection")
	inspectCmd.Flags().Bool("show-kernel-uuid-only", false, "Print only the UUID of the kernel")
	inspectCmd.Flags().Bool("show-kext-load-addresses", false, "Print the load addresses of each kext")
	inspectCmd.Flags().Bool("show-kext-uuids", false, "Print the UUIDs of each kext")
	inspectCmd.Flags().Bool("show-mach-boot-properties", false, "Print derived Mach-O boot properties")
	inspectCmd.Flags().Bool("show-mach-header", false, "Print the mach header(s) in the collection(s)")
	inspectCmd.Flags().Bool("show-prelink-info", false, "Dump the __PRELINK_INFO section(s)")
	inspectCmd.Flags().StringP("system-path", "S", "", "Path to the system kernel collection")
	inspectCmd.Flags().StringP("variant-suffix", "V", "", "Image suffix for the kernel variant")
	inspectCmd.Flags().BoolP("verbose", "v", false, "Toggle verbosity")
	rootCmd.AddCommand(inspectCmd)

	carapace.Gen(inspectCmd).FlagCompletion(carapace.ActionMap{
		"arch":           carapace.ActionValues("arm64", "arm64e", "x86_64", "x86_64h"),
		"auxiliary-path": carapace.ActionFiles(),
		"boot-path":      carapace.ActionFiles(),
		"bundle-path":    carapace.ActionFiles(),
		"system-path":    carapace.ActionFiles(),
	})
}
