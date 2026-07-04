package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showloadedCmd = &cobra.Command{
	Use:   "showloaded",
	Short: "Show the loaded state of extensions on the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showloadedCmd).Standalone()

	showloadedCmd.Flags().StringP("arch", "a", "", "Architecture to use")
	showloadedCmd.Flags().Bool("arch-info", false, "Include architecture info in output")
	showloadedCmd.Flags().StringP("bundle-identifier", "b", "", "Search for and include this identifier")
	showloadedCmd.Flags().StringP("bundle-path", "p", "", "Include the bundle specified at this path")
	showloadedCmd.Flags().String("collection", "", "Restrict load info to a particular collection kind")
	showloadedCmd.Flags().String("elide-identifier", "", "Do not include this identifier in the results")
	showloadedCmd.Flags().StringP("filter", "f", "", "Filter results by predicate (overridden by other arguments)")
	showloadedCmd.Flags().StringP("filter-all", "F", "", "Filter results by predicate (not overridden by other arguments)")
	showloadedCmd.Flags().BoolP("help", "h", false, "Show help information")
	showloadedCmd.Flags().Bool("list-only", false, "Print the list only, omitting column headers")
	showloadedCmd.Flags().BoolP("no-authorization", "z", false, "Skip approval checks")
	showloadedCmd.Flags().Bool("no-kernel-components", false, "Omit kernel components from output")
	showloadedCmd.Flags().String("optional-identifier", "", "Search for and include this identifier if possible")
	showloadedCmd.Flags().Bool("show", false, "Restrict output: loaded, unloaded, or all")
	showloadedCmd.Flags().Bool("show-kernel", false, "Show kernel information in output")
	showloadedCmd.Flags().Bool("show-mach-headers", false, "Show the Mach-O headers of loaded extensions")
	showloadedCmd.Flags().Bool("sort", false, "Sort by load address instead of by load tag")
	showloadedCmd.Flags().StringP("variant-suffix", "V", "", "Image suffix for the kernel variant")
	showloadedCmd.Flags().BoolP("verbose", "v", false, "Toggle verbosity")
	rootCmd.AddCommand(showloadedCmd)

	carapace.Gen(showloadedCmd).FlagCompletion(carapace.ActionMap{
		"arch":        carapace.ActionValues("arm64", "arm64e", "x86_64", "x86_64h"),
		"bundle-path": carapace.ActionFiles(),
		"collection":  carapace.ActionValues("boot", "system", "auxiliary"),
	})
}
