package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unloadCmd = &cobra.Command{
	Use:   "unload",
	Short: "Unload the named kexts and all personalities",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unloadCmd).Standalone()

	unloadCmd.Flags().StringP("arch", "a", "", "Architecture to use")
	unloadCmd.Flags().StringP("bundle-identifier", "b", "", "Search for and include this identifier")
	unloadCmd.Flags().StringP("bundle-path", "p", "", "Include the bundle specified at this path")
	unloadCmd.Flags().StringP("class-name", "c", "", "Terminate all instances of IOService class")
	unloadCmd.Flags().BoolP("help", "h", false, "Show help information")
	unloadCmd.Flags().BoolP("no-authorization", "z", false, "Skip approval checks")
	unloadCmd.Flags().BoolP("personalities-only", "P", false, "Terminate services and remove personalities only")
	unloadCmd.Flags().StringP("variant-suffix", "V", "", "Image suffix for the kernel variant")
	unloadCmd.Flags().BoolP("verbose", "v", false, "Toggle verbosity")
	rootCmd.AddCommand(unloadCmd)

	carapace.Gen(unloadCmd).FlagCompletion(carapace.ActionMap{
		"arch":        carapace.ActionValues("arm64", "arm64e", "x86_64", "x86_64h"),
		"bundle-path": carapace.ActionFiles(),
	})
}
