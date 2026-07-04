package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load one or more extensions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loadCmd).Standalone()

	loadCmd.Flags().StringP("arch", "a", "", "Architecture to use")
	loadCmd.Flags().String("build", "", "System build version number")
	loadCmd.Flags().StringP("bundle-identifier", "b", "", "Search for and include this identifier")
	loadCmd.Flags().StringP("bundle-path", "p", "", "Include the bundle specified at this path")
	loadCmd.Flags().BoolP("help", "h", false, "Show help information")
	loadCmd.Flags().String("kdk", "", "Path to the kdk to use")
	loadCmd.Flags().String("load-style", "", "Load style: start-only, match-only, start-and-match")
	loadCmd.Flags().BoolP("no-authorization", "z", false, "Skip approval checks")
	loadCmd.Flags().BoolP("no-default-repositories", "e", false, "Don't use default repositories for kexts")
	loadCmd.Flags().StringP("personality-name", "P", "", "Send the named personality to the catalog")
	loadCmd.Flags().StringP("repository", "r", "", "Paths to directories containing extensions")
	loadCmd.Flags().String("sdk", "", "Path to the sdk to use")
	loadCmd.Flags().BoolP("verbose", "v", false, "Toggle verbosity")
	loadCmd.Flags().StringP("volume-root", "R", "", "The target volume")
	rootCmd.AddCommand(loadCmd)

	carapace.Gen(loadCmd).FlagCompletion(carapace.ActionMap{
		"arch":        carapace.ActionValues("arm64", "arm64e", "x86_64", "x86_64h"),
		"bundle-path": carapace.ActionFiles(),
		"kdk":         carapace.ActionFiles(),
		"load-style":  carapace.ActionValues("start-only", "match-only", "start-and-match"),
		"repository":  carapace.ActionDirectories(),
		"sdk":         carapace.ActionDirectories(),
		"volume-root": carapace.ActionDirectories(),
	})
}
