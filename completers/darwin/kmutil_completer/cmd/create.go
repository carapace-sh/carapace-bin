package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create one or more new kernel collection artifacts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().StringP("arch", "a", "", "Architecture to use")
	createCmd.Flags().StringP("auxiliary-path", "A", "", "Path to the auxiliary kernel collection")
	createCmd.Flags().StringP("boot-path", "B", "", "Path to the boot kernel collection")
	createCmd.Flags().String("build", "", "System build version number")
	createCmd.Flags().StringP("bundle-identifier", "b", "", "Search for and include this identifier")
	createCmd.Flags().StringP("bundle-path", "p", "", "Include the bundle specified at this path")
	createCmd.Flags().Bool("compress", false, "Compress the collection using LZFSE")
	createCmd.Flags().StringP("filter", "f", "", "Filter results by predicate (overridden)")
	createCmd.Flags().StringP("filter-all", "F", "", "Filter results by predicate (not overridden)")
	createCmd.Flags().Bool("force", false, "Force a collection rebuild even if nothing changed")
	createCmd.Flags().BoolP("help", "h", false, "Show help information")
	createCmd.Flags().Bool("img4-encode", false, "Encode the collection in an img4 payload")
	createCmd.Flags().Bool("include-volume-kexts", false, "Include kexts installed on the ARV (Apple Silicon only)")
	createCmd.Flags().String("kdk", "", "Path to the kdk to use")
	createCmd.Flags().StringP("kernel", "k", "", "Path to the kernel to use")
	createCmd.Flags().StringP("manifest", "m", "", "Source build manifest for kernel collection rebuilding")
	createCmd.Flags().StringP("new", "n", "", "Which types of artifact to build: boot, sys, or aux")
	createCmd.Flags().BoolP("no-authorization", "z", false, "Skip approval checks")
	createCmd.Flags().BoolP("no-system-collection", "L", false, "Don't look for or generate a system collection")
	createCmd.Flags().Bool("no-variant-extension", false, "Don't append build variant as extension to path")
	createCmd.Flags().StringP("repository", "r", "", "Paths to directories containing extensions")
	createCmd.Flags().String("sdk", "", "Path to the sdk to use")
	createCmd.Flags().Bool("split-ro-rx-tags", false, "Split RO RX region payload property tags for iBoot")
	createCmd.Flags().StringP("strip", "s", "", "Strip mode: none, all, or allkexts")
	createCmd.Flags().StringP("system-path", "S", "", "Path to the system kernel collection")
	createCmd.Flags().StringP("variant-suffix", "V", "", "Image suffix for the kernel variant")
	createCmd.Flags().BoolP("verbose", "v", false, "Toggle verbosity")
	createCmd.Flags().StringP("volume-root", "R", "", "The target volume")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"arch":           carapace.ActionValues("arm64", "arm64e", "x86_64", "x86_64h"),
		"auxiliary-path": carapace.ActionFiles(),
		"boot-path":      carapace.ActionFiles(),
		"bundle-path":    carapace.ActionFiles(),
		"kdk":            carapace.ActionFiles(),
		"kernel":         carapace.ActionFiles(),
		"manifest":       carapace.ActionFiles(),
		"new":            carapace.ActionValues("boot", "sys", "aux"),
		"repository":     carapace.ActionDirectories(),
		"sdk":            carapace.ActionDirectories(),
		"strip":          carapace.ActionValues("none", "all", "allkexts"),
		"system-path":    carapace.ActionFiles(),
		"volume-root":    carapace.ActionDirectories(),
	})
}
